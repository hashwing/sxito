package db

func AddGateway(gateway *DeviceGateway)error{
	_,err:=MysqlDB.Table("sxito_gateway").Insert(gateway)
	return err
}

func FindGateways(adminID string)([]DeviceGateway,error){
	var gateways []DeviceGateway
	err:=MysqlDB.Table("sxito_gateway").Where("admin_id=?",adminID).Find(&gateways)
	return gateways,err
}

func GetGateway(gatewayID string)(*DeviceGateway,error){
	var gateway DeviceGateway
	res,err:=MysqlDB.Table("sxito_gateway").Where("gateway_id=?",gatewayID).Get(&gateway)
	if !res{
		return nil,nil
	}
	return &gateway,err
}


func UpdateGateway(gateway *DeviceGateway)error{
	_,err:=MysqlDB.Table("sxito_gateway").Where("gateway_id=?",gateway.ID).Update(gateway)
	return err
}

func DelGateway(gatewayID string)error{
	gateway:=new(DeviceGateway) 
	_,err:=MysqlDB.Table("sxito_gateway").Where("gateway_id=?",gatewayID).Delete(gateway)
	return err
}