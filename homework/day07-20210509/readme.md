1. 整理mm图
2. 存储添加json, 支持json，csv，gob => 使用接口方式 不要再函数中判断json，csv，gob去调用不同功能（可通过判断设置接口类型为csv，json，gob）
    通过工厂创建persistent对象
    Factory.Create(type_ string) Persistent
