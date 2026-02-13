package bar

import "conceitos_avancados/foo"



func TakeFoo(i foo.Interface){
	i.Interface()
}