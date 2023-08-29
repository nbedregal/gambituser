package models

type SecretRDSJson struct {
	Username             string `json:"username"`
	Password             string `json:"password"`
	Engine               string `json:"engine"`
	Host                 string `json:"host"`
	Port                 int    `json:"port"`
	Dbname               string `json:"dbname"`
	DbInstanceIdentifier string `json:"dbInstanceIdentifier"`
}

type SignUp struct {
	UserEmail string `json:"UserEmail"`
	UserUUID  string `json:"UserUUID"`
}
