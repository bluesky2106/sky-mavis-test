package api

func (s *Server) Routes() {
	router := s.g
	router.GET("/", s.DefaultWelcome)
	// health and live check
	router.GET("/health", s.healthCheck)
	router.GET("/live", s.liveCheck)

	api := router.Group("/api/v1")
	{
		api.GET("/", s.Welcome)

		// visitor
		visitor := api.Group("/visitors")
		{
			visitor.GET("/current", s.GetCurrentVisitor)
			visitor.GET("/top", s.GetTop100Visitors)
			visitor.GET("/last", s.GetLast100Visitors)
		}
	}
}
