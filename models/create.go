package model
func CreateTodo() error{
	insertQ, err:=con.Query("INSERT INTO TODO VALUES($1,$2)","Daniel","Work1")//insertQ->insert query
defer insertQ.Close()
if err !=nil{return err}
return nil
}