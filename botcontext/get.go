package botcontext

import (
	"fmt"

	"github.com/jadevelopmentgrp/Tickets-Utilities/restcache"

	"github.com/jadevelopmentgrp/Tickets-Import-API/config"
	"github.com/jadevelopmentgrp/Tickets-Import-API/redis"
	"github.com/rxdn/gdl/rest/ratelimit"
)

func ContextForGuild(guildId uint64) (*BotContext, error) {
	rateLimiter := ratelimit.NewRateLimiter(ratelimit.NewRedisStore(redis.Client.Client, fmt.Sprintf("ratelimiter:%d", config.Conf.Bot.Id)), 1)

	return &BotContext{
		BotId:       config.Conf.Bot.Id,
		Token:       config.Conf.Bot.Token,
		RateLimiter: rateLimiter,
		RestCache:   restcache.NewRedisRestCache(redis.Client.Client, config.Conf.Bot.Token, rateLimiter),
	}, nil
}
