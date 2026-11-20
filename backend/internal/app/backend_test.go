package app

import "github.com/Keramikus-97/omastore/backend/internal/rpc"

// *App must satisfy rpc.Backend.
var _ rpc.Backend = (*App)(nil)
