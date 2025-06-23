package handler


func GetEmployeeByIDHandler(c *gin.Context){

	employee, err := repository.GetEmplyeeByID(employeeID)
	if err != nil{
		handleEmployeeError(c,err)
	}

	c.JSON(http.StatusOK, employee)
}

func GetEmployeeByRegistration(c *gin.Context){
	registration := c.Param("registration")

	employee, err := repository.GetEmployeeByRegistration(registration)
	if err != nil{
		handleEmployeeError(c, err)
	}
	c.JSON(http.StatusOK, employee)
}

func handleEmployeeError(c *gin.Context, err error){
	if errors.Is(err, sql.ErrNoRows){
		c.JSON(http.StatusNotFound, gin.H{"error": "Funcionário não encontrado"})
	}else{
		c.JSON(http.StatusInternalServeError, gin.H{
			"error":"Erro ao buscar funcionário",
			"details": err.Error(),
		})
	}
}