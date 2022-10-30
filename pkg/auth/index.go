package auth

const (
	HASH_COST = 10
)

// type jwtClaims struct {
// 	Name    string `json:"name"`
// 	UUID    string `json:"uuid"`
// 	IsAdmin bool   `json:"admin"`
// 	jwt.StandardClaims
// }

// func AddAuthActions(echoServer *echo.Echo) {
// 	authRoutes := echoServer.Group("/auth/true")

// 	jwtConfig := middleware.JWTConfig{
// 		Claims:     &jwtClaims{},
// 		SigningKey: []byte(os.Getenv(cc.SIGNING_KEY)),
// 	}

// 	authRoutes.Use(middleware.JWTWithConfig(jwtConfig))

// 	authRoutes.GET("/", func(c echo.Context) error {
// 		return c.String(http.StatusOK, "hella lit")
// 	})
// }
