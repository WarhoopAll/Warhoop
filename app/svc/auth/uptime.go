package auth

import (
	"context"
	"net"
	"strconv"
	"time"
	"warhoop/app/config"
	"warhoop/app/model/auth"
	"warhoop/app/utils"
)

func (s *AuthService) GetUptime(ctx context.Context) (*auth.Uptime, error) {
	entry, err := s.store.AuthRepo.GetUptimeByID(ctx, config.Get().RealmID)
	if err != nil {
		return nil, utils.ErrDataBase
	}

	if !s.CheckOnlineStatus(config.Get().RealmRealmlist, entry.Realm.Port) {
		return nil, utils.ErrInternal
	}

	return entry.ToWeb(), nil
}

func (s *AuthService) CheckOnlineStatus(host string, port int16) bool {
	address := net.JoinHostPort(host, strconv.Itoa(int(port)))

	timeout := 1 * time.Second

	conn, err := net.DialTimeout("tcp", address, timeout)
	if err != nil {
		return false
	}
	conn.Close()

	return true
}
