package domain

/*
func TestSearchProductHappy(t *testing.T) {

	//Arrange
	//mockRepo := &product.MockRepository{}
	testProdcuts := []*Product{
		{1, "Oil - Margarine", 439, "S82254D", true, "15/12/2021", 71.42},
		{2, "Pineapple - Canned, Rings", 345, "M4637", true, "09/08/2021", 352.79},
		{3, "Wine - Red Oakridge Merlot", 367, "T65812", false, "24/05/2021", 179.23},
		{4, "Cookie - Oatmeal", 130, "M7157", false, "28/01/2022", 275.47},
		{5, "Flavouring Vanilla Artificial", 336, "S60152S", true, "10/02/2022", 839.02},
		{6, "Cake - Lemon Chiffon", 446, "S51821A", true, "06/04/2022", 895.88},
	}
	//esto no se si esta bien (?)
	mockRepo.Products = testProdcuts
	//sp := product.NewSuperMarket(mockRepo)

	expectRes := []Product{
		{5, "Flavouring Vanilla Artificial", 336, "S60152S", true, "10/02/2022", 839.02},
		{6, "Cake - Lemon Chiffon", 446, "S51821A", true, "06/04/2022", 895.88},
	}
	//Act
	actualRes, err := sp.SearchProduct("801.0")
	//Assert
	if actualRes[0].Id != expectRes[0].Id {
		t.Errorf("something went wrong expecetd result is %d but got %d", actualRes[0].Id, expectRes[0].Id)
	}
	if err != nil || len(actualRes) != 2 {
		t.Errorf("something went wrong expecetd result is %d but got %d", 2, len(actualRes))
	}
}

func TestSearchProductsNull(t *testing.T) {
	//Arrange
	mockRepo := &product.MockRepository{}
	sp := product.NewSuperMarket(mockRepo)
	//Act
	_, err := sp.SearchProduct("801.0")
	//Assert
	if err == nil {
		t.Errorf("something went wrong this test should return an error ")
	}
}
*/
