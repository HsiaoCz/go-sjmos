## go设计模式

软件的开发学习可以当作武功修炼的话，可以分为招式和内功

内功:
数据结构
算法
设计模式
架构设计
软件工程

招式很快就能学会，但是内功修炼需要很长时间

模式：每个模式都描述了一个在我们环境中不断出现的问题，然后描述了该问题的解决方案的核心
通过这种方式我们可以无数次的重复使用那些已有的成功的额解决方案，无须再重复相同的工作

简单说，在一定的环境下，用固定套路解决问题

设计模式的种类
(1):创建型模式:如何创建对象
(2):结构型模式:如何实现类或对象的组合
(3):行为型模式:类或对象 怎样交互以及怎样分配职责

简单工厂模式不属于GoF 23种设计模式

### 软件设计模式的作用

思考这样几个问题:
1.如何将代码分散在几个不同的类中？
2.为什么要有接口？
3.何谓针对抽象编程
4.何时不应该使用继承
5.如何不修改源代码增加新功能？
6.更好地阅读和理解现有类库与其他系统中的源代码



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
