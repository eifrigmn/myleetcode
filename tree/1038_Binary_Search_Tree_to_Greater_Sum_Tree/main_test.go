package tree 
import (
	"myleetcode/util"
	"reflect"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)
func TestBstToGst(t *testing.T) {
	Convey("从二叉搜索树到更大和树", t, func(){
		//        4								30
		//     /     \						/       \
		//    /       \					   /         \
		//    1        6                  36         21
		//  /   \     /  \        ===>    /         /  \ 
		//  0   2    5    7              36        26  15
		//       \         \                             \
		//        3         8                             8
		var null = util.NULL
		nums := []int{4,1,6,0,2,5,7,null,null,null,3,null,null,null,8}
		tree := util.IntsToTreeNode(nums)
		res := bstToGst(tree)
		resArry := util.TreeToInts(res)
		So(reflect.DeepEqual(resArry, []int{30,36,21,36,35,26,15,null,null,null,33,null,null,null,8}), ShouldBeTrue)
	})
}