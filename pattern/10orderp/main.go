package main

import "fmt"

// 如果不用命令模式
// doctor命令的接收者
type Docter struct {
}

func (d *Docter) treateEye() {
	fmt.Println("治疗眼睛...")
}

func (d *Docter) treateNose() {
	fmt.Println("治疗鼻子....")
}

// 病人 业务层

func main() {
	// 病人填写病单 病人依赖病单看病
	docter := new(Docter)
	cmdEye := CommandTreatEye{docter: docter}
	// 看鼻子
	cmdNose := CommandTreatNose{docter: docter}

	// 护士
	nurse := new(Nurse)
	// 收集管理病单
	nurse.CmdList = append(nurse.CmdList, &cmdEye, &cmdNose)

	// 执行病单指令
	nurse.Notify()
}

// 使用命令模式
// 治疗眼睛的病单
type CommandTreatEye struct {
	docter *Docter
}

func (c *CommandTreatEye) Treat() {
	c.docter.treateEye()
}

// 治疗鼻子的病单
type CommandTreatNose struct {
	docter *Docter
}

func (ct *CommandTreatNose) Treat() {
	ct.docter.treateNose()
}

// 抽象的命令
type Command interface {
	Treat()
}

// 护士，命令的调用者
type Nurse struct {
	CmdList []Command //收集命令
}

// 护士发送命令的方法
func (n *Nurse) Notify() {
	if n.CmdList == nil {
		return
	}
	for _, cmd := range n.CmdList {
		cmd.Treat() //多态现象，调用具体的命令
		// 就会调用病单已经绑定的医生的诊断方法
	}
}
