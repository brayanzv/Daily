package models

/*GraboTweet es la escructura usada para usar el Tweet*/
type Tweet struct {
	Mensaje string `bson:"mensaje" json: "mensaje"`
}
