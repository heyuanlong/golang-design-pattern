# 桥接模式

桥接模式分离抽象部分和实现部分。使得两部分独立扩展。

桥接模式类似于策略模式，区别在于策略模式封装一系列算法使得算法可以互相替换。

策略模式使抽象部分和实现部分分离，可以独立变化。


桥接（Bridge）模式包含以下主要角色。
    抽象化（Abstraction）角色：定义抽象类，并包含一个对实现化对象的引用。
    扩展抽象化（Refined Abstraction）角色：是抽象化角色的子类，实现父类中的业务方法，并通过组合关系调用实现化角色中的业务方法。
    实现化（Implementor）角色：定义实现化角色的接口，供扩展抽象化角色调用。
    具体实现化（Concrete Implementor）角色：给出实现化角色接口的具体实现。