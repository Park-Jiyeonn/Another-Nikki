package service

import (
	"github.com/google/wire"
)

// ProviderSet is service providers.
var ProviderSet = wire.NewSet(
	NewProblemService,
	NewUserService,
	NewArticleService,
	NewCommentService,
	NewCodeProcessingService,
	NewLogsService,
)
