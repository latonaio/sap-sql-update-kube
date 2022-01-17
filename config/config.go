package config

type Conf struct {
	RMQ *RMQ
	SAP *SAP
	DB  *Database
}

func NewConf() *Conf {
	return &Conf{
		RMQ: newRMQ(),
		SAP: newSAP(),
		DB:  newDatabase(),
	}
}
