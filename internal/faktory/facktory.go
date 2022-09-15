package faktory

import (
	"fmt"
	"path/filepath"

	faktory "github.com/contribsys/faktory/server"
	"github.com/contribsys/faktory/storage"
)

func LocalServer(bind string, dir string) (server *faktory.Server, cancel func(), err error) {
	sock := filepath.Join(dir, "redis.sock")
	cancel, err = storage.Boot(dir, sock)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to start Faktory redis server: %w", err)
	}
	server, err = faktory.NewServer(&faktory.ServerOptions{
		Binding:          bind,
		StorageDirectory: dir,
		RedisSock:        sock,
		// ConfigDirectory:  "",
		// Environment:      "",
		// Password:         "",
		PoolSize: 50,
		//GlobalConfig: map[string]interface{}{
		//	"": nil,
		//},
	})
	if err != nil {
		return nil, nil, fmt.Errorf("failed to create Faktory server: %w", err)
	}
	if err := server.Boot(); err != nil {
		return nil, nil, fmt.Errorf("failed to boot Faktory server: %w", err)
	}
	return server, cancel, nil
}
