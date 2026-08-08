package user

import (
	"testing"

	"xygo/internal/model/entity"
)

func dept(id, parent uint64, name string) *entity.AdminDept {
	return &entity.AdminDept{Id: id, ParentId: parent, Name: name}
}

func user(id, deptId uint64, name string) *entity.AdminUser {
	return &entity.AdminUser{Id: id, DeptId: deptId, RealName: name}
}

// TestBuildUserDeptTree_Basic 基础：用户挂到部门，空部门剪除
func TestBuildUserDeptTree_Basic(t *testing.T) {
	depts := []*entity.AdminDept{
		dept(1, 0, "总部"),
		dept(2, 1, "研发部"),
		dept(3, 1, "市场部"),
		dept(4, 2, "研发一组"),
	}
	users := []*entity.AdminUser{
		user(101, 4, "张三"),
		user(102, 2, "李四"),
	}
	roots := buildUserDeptTree(users, depts)
	if len(roots) != 1 {
		t.Fatalf("期望1个根部门, 实际 %d", len(roots))
	}
	root := roots[0]
	if root.Value != 1 {
		t.Fatalf("根应为总部(1), 实际 %d", root.Value)
	}
	// 市场部(3)无用户应被剪除 → 总部children应有 研发部(2)
	if len(root.Children) != 1 {
		t.Fatalf("总部应有1个子部门, 实际 %d", len(root.Children))
	}
}

// TestBuildUserDeptTree_UserIdCollision 用户ID与部门ID相同不误判
func TestBuildUserDeptTree_UserIdCollision(t *testing.T) {
	depts := []*entity.AdminDept{
		dept(5, 0, "部门A"),
		dept(6, 5, "部门B"),
	}
	// 用户 id=6 恰好等于部门B的id
	users := []*entity.AdminUser{
		user(6, 5, "老王"),
	}
	roots := buildUserDeptTree(users, depts)
	if len(roots) != 1 {
		t.Fatalf("期望1个根, 实际 %d", len(roots))
	}
	root := roots[0]
	// 部门B(6)无用户应被剪除，老王的叶子应挂在部门A下
	if len(root.Children) != 1 {
		t.Fatalf("部门A应有1个用户叶子, 实际 %d", len(root.Children))
	}
	if root.Children[0].Value != 6 || root.Children[0].Label != "老王" {
		t.Fatalf("叶子应为用户6(老王), 实际 %d/%s", root.Children[0].Value, root.Children[0].Label)
	}
}

// TestBuildUserDeptTree_Cycle 部门树成环不死循环
func TestBuildUserDeptTree_Cycle(t *testing.T) {
	depts := []*entity.AdminDept{
		dept(1, 2, "甲"), // 1 -> 2
		dept(2, 1, "乙"), // 2 -> 1 成环
	}
	users := []*entity.AdminUser{
		user(101, 1, "用户1"),
	}
	roots := buildUserDeptTree(users, depts)
	// 不崩溃即可，环中的部门可能异常但不应死循环
	_ = roots
}

// TestBuildUserDeptTree_Orphan 无部门用户作为根
func TestBuildUserDeptTree_Orphan(t *testing.T) {
	depts := []*entity.AdminDept{
		dept(1, 0, "部门A"),
	}
	users := []*entity.AdminUser{
		user(201, 0, "游离用户"),
		user(202, 1, "部门用户"),
	}
	roots := buildUserDeptTree(users, depts)
	if len(roots) != 2 {
		t.Fatalf("期望2个根(部门A + 游离用户), 实际 %d", len(roots))
	}
}
