type handler struct {  
	store datastore.Student  
  }  
	
  func New(s datastore.Student) handler {  
	return handler{store: s}  
  }  
	
  func (h handler) GetByID(ctx *gofr.Context) (interface{}, error) {  
	// ctx.PathParam() returns the path parameter from HTTP request.
	id := ctx.PathParam("id")  
	if id == "" {  
	return nil, errors.MissingParam{Param: []string{"id"}}  
	}  
	
	if _, err := validateID(id); err != nil {  
	return nil, errors.InvalidParam{Param: []string{"id"}}  
	}  
	
	resp, err := h.store.GetByID(ctx, id)  
	if err != nil {  
	return nil, errors.EntityNotFound{  
   Entity: "student",  
	ID:     id,  
	}  
	}  
	
	return resp, nil  
  }
  
  func validateID(id string) (int, error) {  
	res, err := strconv.Atoi(id)  
	if err != nil {  
	return 0, err  
	}  
	
	return res, err  
  }