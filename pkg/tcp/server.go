package tcp

import (
	"crypto/tls"
	"fmt"
	"time"

	"go.osspkg.com/goppy/errors"
	"go.osspkg.com/goppy/iosync"
	"go.osspkg.com/goppy/xc"
)

type Server struct {
	conf      TCPConfig
	listeners []*Listen
	wg        iosync.Group
}

func New(conf TCPConfig) *Server {
	return &Server{
		conf: conf,
		wg:   iosync.NewGroup(),
	}
}

func (v *Server) Up(ctx xc.Context) error {
	if err := v.buildListeners(); err != nil {
		return err
	}

	for _, l := range v.listeners {
		l := l
		v.wg.Background(func() {
			for {
				conn, err := l.Accept()
				if err != nil {
					select {
					case <-ctx.Done():
						return
					default:
					}
					fmt.Println(err.Error())
					time.Sleep(time.Second * 1)
					continue
				}

				if v.conf.Timeout > 0 {
					conn.SetDeadline(time.Now().Add(v.conf.Timeout))
				}

				if tc, ok := conn.(*tls.Conn); ok {
					if err = tc.HandshakeContext(ctx.Context()); err != nil {
						fmt.Println(err.Error(), conn.Close())
						continue
					}
				}

				v.wg.Background(func() {

				})
			}
		})
	}

	return nil
}

func (v *Server) Down() error {
	var err error
	for _, l := range v.listeners {
		err = errors.Wrap(err, l.Close())
	}
	v.wg.Wait()
	return err
}

func (v *Server) buildListeners() error {
	for _, c := range v.conf.Config {
		certs := make([]Cert, 0, len(c.Certs))
		for _, cert := range c.Certs {
			certs = append(certs, Cert{Public: cert.Public, Private: cert.Public})
		}
		l, err := NewListen(c.Port, certs...)
		if err != nil {
			return err
		}
		v.listeners = append(v.listeners, l)
	}
	return nil
}
