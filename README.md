## go 设计模式

软件的开发学习可以当作武功修炼的话，可以分为招式和内功

内功:
数据结构
算法
设计模式
架构设计
软件工程

招式很快就能学会，但是内功修炼需要很长时间

模式：每个模式都描述了一个在我们环境中不断出现的问题，然后描述了该问题的解决方案的核心
通过这种方式我们可以无数次的重复使用那些已有的成功的解决方案，无须再重复相同的工作

简单说，在一定的环境下，用固定套路解决问题

设计模式的种类
(1):创建型模式:如何创建对象
(2):结构型模式:如何实现类或对象的组合
(3):行为型模式:类或对象 怎样交互以及怎样分配职责

简单工厂模式不属于 GoF 23 种设计模式

### 软件设计模式的作用

思考这样几个问题: 1.如何将代码分散在几个不同的类中？ 2.为什么要有接口？ 3.何谓针对抽象编程 4.何时不应该使用继承 5.如何不修改源代码增加新功能？ 6.更好地阅读和理解现有类库与其他系统中的源代码

## 1、设计原则

1.单一职责原则(single responsibility principle)

类的职责单一，对外只提供一种功能，而引起类变化的原因都应该只有一个
单一职责原则，类的职责尽量单一
但是项目很复杂的时候，类的职责单一会导致有许多的类，类与类之间的耦合度会增加

2.开闭原则(open-closed)

类的改动是通过增加代码进行的，而不是修改代码

3、里氏代换原则(liskov Substitution principle)

任何抽象类(interface)出现的地方都可以用他的实现类来进行替换，实际就是虚拟机制，语言级别实现面向对象功能

4、依赖倒转原则(Dependence inversion Principle)

依赖于抽象，而不是依赖于具体的类，也是针对抽象编程

5、接口隔离(interface Segregation principle)

不应该强迫用户的程序依赖它们不需要的接口方法。一个接口应该只提供一种对外的功能，不应该把所有操作都封装到一个接口中去

6、合成复用原则(Composite Reuse principle)

如果使用继承，会导致父类的任何变换都可能影响到子类的行为。如果使用对象组合，就降低了这种依赖关系。对于继承和组合，优先使用组合

7、迪米特法则(Law of Demeter,LoD)

一个对象应当对其他对象尽可能少的了解，从而降低各个对象之间的耦合，提高系统的可维护性
例如在一个程序中，各个模块之间相互调用时，通常会提供一个统一的接口来实现。
这样其他模块不需要了解另外一个模块的内部实现细节，这样当一个模块内部实现发生改变时，
不会影响其他模块的使用。

迪米特法则又叫知道最少原则

举个例子 如果一个人要买房子
他可能需要将市场上所有的房子都看一遍
如果另外一个人也需要买房 他也需要去实地了解所有的房子的信息

这时候整个系统的耦合度就很高了

这时候来一个中介 每一个买房的人都只从中介那里了解房子的信息
房子的信息 只与中介耦合

就是添加中间层

使两边只与中间层耦合

## 设计模式

### 1、创建型模式

单例模式:保证一个类只有一个实例，并提供一个访问它的全局访问点
简单工厂模式:通过专门定义一个类来负责创建其他类的实例，被创建的实例通常都具有共同的父类
工厂方法模式:定义一个创建产品对象的工厂接口，将实际创建工作推迟到子类中
抽象工厂模式:提供一个创建一系列相关或者互相依赖的接口，而无需指定它们具体的类
原型模式:用原型实例指定创建的对象的种类，并且通过拷贝这些原型创建新的对象
建造者模式：将一个复杂的构建与其表示相分离，使得同样的构建过程可以创建不同的表示

### 1、简单工厂模式

首先:为什么需要简单工厂模式

简单工厂模式:
工厂角色:简单工厂的核心，它负责创建所有实例的内部逻辑。工厂类可以被外界直接调用，创建所需的产品对象
抽象产品角色：简单工厂所创建的所有对象的父类，它负责描述所有实例所共有的公共接口
具体产品角色：简单工厂模式所创建的具体实例对象

简单工厂方法模式的优点： 1.实现了对象创建和使用的分离 2.不需要记住具体的类名，记住参数即可，减少使用者的记忆量

缺点： 1.对工厂类职责过重，一旦不能工作，系统受到影响 2.增加系统中类的个数，复杂度和理解度增加 3.违反了开闭原则，增加产品需要修改工厂逻辑，工厂越来越复杂

适用场景：
工厂类负责创建对象比较少，由于创建的对象较少，不会造成工厂方法中的业务逻辑太过复杂
客户端只知道传入工厂的参数，对于如何创建对象并不关心

### 2、工厂模式

简单工厂方法模式 是不符合开闭原则的
工厂方法模式可以看作是简单工厂方法模式加上开闭原则

使用工厂方法模式，类有类的抽象接口，工厂有工厂的抽象接口
每一个类有一个工厂类，每一个工厂类都实现了抽象工厂类

工厂方法模式的扩展性更好
如果我们需要添加一个类
我可以在不影响别的类的基础上扩张

工厂方法的缺点在于 增加了类的数量

### 3、抽象工厂方法模式

工厂方法模式每添加一个类对应的就需要添加一个工厂类
抽象工厂方法模式可以解决这个问题

在了解抽象方法模式之前先了解产品族与产品等级结构
将一个系列的产品叫一个产品族
比如中国产品族，日本产品族，美国产品族
所以工厂就按照产品族来划分

抽象工厂方法在产品固定的时候，是符合开闭的
但一旦对产品结构进行扩展，就完全不符合开闭原则了

### 4、单例模式

单例模式要解决的问题是：
保证一个类永远只能有一个对象，且该对象的功能依然能被其他模块使用
单例模式的饿汉式：在没有获取单例的时候就已经分配内存了
单例模式的懒汉式：在获取的时候才分配内存

```go
var instance *singleton

func GetInstance()*singleton{
    // 只有首次GetInstance方法被调用才会产生单例对象
    if instance==nil{
        instance=new(singleton)
        return instance
    }
    return instance
}

// 但这样会存在一个问题 并发不安全 ，当有两个goroutine同时访问的时候，会造成内存泄漏
// 一种方法是加锁
// 但是加锁又会造成性能问题
// 可以考虑使用原子操作
```

```go
var instanced uint32

func GetInstance()*singleton{
  if atomic.LoadUint32(&instanced==1){
    return instance
  }
  return instance
}
```

当然还可以使用 once
保证一个操作只执行一次

```go
var once sync.Once

func GetInstance()*singleton{
    once.Do(func{
        instance=new(singleton)
    })
}
```

单例模式什么时候来使用
一个系统中只能有一个实例可以使用单例

## 结构型模式

1、适配器模式
将一个类的接口转换成客户希望的另外一个接口使得原本由于接口不兼容而不能一起工作的那些类可以一起工作

2、桥接模式
将抽象部分于实际部分分离，使得它们可以独立的变化

3、组合模式
将对象组合成树形结构以表示"部分-整体"的层次结构，使得用户对单个对象的使用具有一致性

4、装饰模式
动态的给一个对象添加一些额外的职责，就增加功能来说，次模式比生成子类更为灵活

5、外观模式
为子系统中的一组接口提供一个一致的界面，此模式定义了一个高层接口，这个接口使得这个子系统更加容易使用

6、享元模式
以共享的方式高效的支持大量的细粒度的对象

7、代理模式
为其他对象提供一种代理以控制对这个对象的访问

### 1、代理模式

假如说有一个网络游戏 我们自己需要给游戏角色升级，但是我自己比较懒，可以找一个专业的代理替自己给游戏角色升级
专业代练不仅能升级，还有附加收益

代理模式的优缺点：

优点：
能够协调调用者和被调用者，在一定程度上降低了系统的耦合度。
客户端可以针对主题角色进行编程，增加和更换代理类无需修改源代码，符合开闭原则，系统具有较好的灵活性和可扩展性

缺点也很明显
代理的实现较为复杂

### 2、代理模式和谐版

```go
// 代理模式和谐版

// 抽象的主题
type BeautyWoman interface {
	//对男人抛媚眼
	MakeEyesWithMan()
	//和男人共度良宵
	HappyWithMan()
}

// 具体的主题
// panjinlian
type PanJinlian struct{}

// 抛媚眼
func (p *PanJinlian) MakeEyesWithMan() { fmt.Println("给西门大官人抛了个媚眼") }

// 和男人共度美好的时刻
func (p *PanJinlian) HappyWithMan() { fmt.Println("和西门大官人共度了美好的时光...") }

// 中间的代理人
type WangPo struct {
	woman BeautyWoman
}

func NewProxy(woman BeautyWoman) BeautyWoman {
	return &WangPo{woman: woman}
}

// 抛媚眼
func (w *WangPo) MakeEyesWithMan() { w.woman.MakeEyesWithMan() }

// 和男人共度美好的时刻
func (w *WangPo) HappyWithMan() { w.woman.HappyWithMan() }

// 业务逻辑 西门大官人

func main() {
	// 大官人想找金莲，让王婆来安排
	wangpo := NewProxy(new(PanJinlian))
	// 王婆命令金莲
	wangpo.MakeEyesWithMan()
	// 王婆带着大官人
	wangpo.HappyWithMan()
}
```

### 3、装饰器模式

比如一个手机，经过一个贴膜装饰器类，变成了一个贴膜的手机，一个新的类
也可以经过一个装手机壳的装饰器类，变成一个戴手机壳的手机

### 4、适配器模式

将一个类的接口转换成客户希望的另外一个接口。使得原本由于接口不兼容而不能一起工作的那些类可以一起工作

这里面的典型场景，比如手机的充电器，你像phone的充电电压只有5v,可以家用的电压220v,充电器就可以将220v转成5v

适配器模式的优缺点
可以将类和适配者类解耦
增加类的透明性和复用性
灵活扩展性好

缺点：适配器中置换适配者类的某些方法比较麻烦

适用场景:系统需要适用一些现有的类，而这些类的接口不符合系统的需求

### 5、外观模式

在家里可能有多个电器，每个电器有一个遥控器
想要使用每个电器，就需要使用各自的遥控器

我们可以使用一个路由器来连接不同地家电，我们只需要使用智能遥控器，遥控智能路由器
就能遥控整个家电
而这个智能路由器就可以看作是外观

优点是对客户端屏蔽了系统组件，减少了客户端所需要处理地对象的数目
使得子系统更加容易使用，通过引入外观模式，客户端代码将变得很简单
实现了子系统和客户端之间地松耦合，子系统的变化只需要调整外观
一个子系统地修改对其他子系统没有影响

缺点：
不能很好地限制客户端直接使用子系统类，如果客户端访问子系统做太多地限制则减少了可变性和灵活性
如果设计不当，增加新的子系统可能需要修改外观类的源代码，违背了开闭原则

适用于复杂系统需要简单入口使用
或者客户端程序和多个子系统之间存在很大的依赖性
