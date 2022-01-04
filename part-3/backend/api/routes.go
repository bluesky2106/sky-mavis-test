package api

func (s *Server) Routes() {
	router := s.g
	router.GET("/", s.DefaultWelcome)
	// health and live check
	router.GET("/health", s.healthCheck)
	router.GET("/live", s.liveCheck)
}
