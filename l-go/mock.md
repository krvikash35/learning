mock tool
* https://github.com/uber-go/mock
* https://github.com/golang/mock -> called `gomock` or `mockgen`, deprecated by above uber
    ```
    mockgen -destination=./mocks.go -package=voucher -mock_names=API=APIMock . API
    ```

    ```
    type Foo interface {
        Bar(x int) int
    }

    func SUT(f Foo) {
        // ...
    }

    func TestFoo(t *testing.T) {
        ctrl := gomock.NewController(t)
        m := NewMockFoo(ctrl)
        m.EXPECT().Bar(gomock.Eq(99)).Return(101)
        SUT(m)
    }
    ```
* https://github.com/vektra/mockery
    * mockery provides the ability to easily generate mocks for Golang interfaces using below `stretchr/testify/mock` package. It removes the boilerplate coding required to use mocks.
    ```
    mockery --dir=. --name=API --filename=mocks.go --output=. --outpkg=voucher
    ```

    ```
    type DB interface {
        Get(val string) string
    }

    func getFromDB(db DB) string {
        return db.Get("ice cream")
    }
    ```
    ```
    import (
        "testing"
        "github.com/stretchr/testify/assert"
    )

    func Test_getFromDB(t *testing.T) {
        mockDB := NewMockDB(t)
        mockDB.EXPECT().Get("ice cream").Return("chocolate").Once()
        flavor := getFromDB(mockDB)
        assert.Equal(t, "chocolate", flavor)
    }

    ```

* https://github.com/stretchr/testify
    * it provide `assert` and `suite` pkg for setup & teardown
    ```
    import (
        "testing"
        "github.com/stretchr/testify/assert"
        "github.com/stretchr/testify/suite"
    )

    type ExampleTestSuite struct {
        suite.Suite
        VariableThatShouldStartAtFive int
    }

    func (suite *ExampleTestSuite) SetupTest() {
        suite.VariableThatShouldStartAtFive = 5
    }

    func (suite *ExampleTestSuite) TestExample() {
        assert.Equal(suite.T(), 5, suite.VariableThatShouldStartAtFive)
    }

    func TestExampleTestSuite(t *testing.T) {
        suite.Run(t, new(ExampleTestSuite))
    }
    ```
    * it also provide pkg to mock at `github.com/stretchr/testify/mock`
    * above tool called vektra's `mockery` can be used to auto generate mock

    ```
    package yours

    import (
        "testing"
        "github.com/stretchr/testify/mock"
    )

    type MyMockedObject struct{
        mock.Mock
    }

    func (m *MyMockedObject) DoSomething(number int) (bool, error) {
        args := m.Called(number)
        return args.Bool(0), args.Error(1)
    }

    func TestSomething(t *testing.T) {
        testObj := new(MyMockedObject)
        testObj.On("DoSomething", 123).Return(true, nil)
        targetFuncThatDoesSomethingWithObj(testObj)
        testObj.AssertExpectations(t)
    }
    ```
    


