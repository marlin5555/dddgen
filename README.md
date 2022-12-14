# dddgen
![Coverage](https://img.shields.io/badge/Coverage-30.2%25-yellow)

[![License: GPL](https://img.shields.io/badge/License-GNU%20GPL-blue)](https://github.com/marlin5555/dddgen/blob/main/LICENSE)

DDD（domain driven design）分层模型的代码生成器。

## 为什么要做代码生成器？
在工程应用中，绝大部分的编码工作就是按照业务逻辑进行技术实现。
这里面包含两部分工作：1. 对应用涉及的领域知识进行建模；2. 使用良好的技术达成领域逻辑。
实际项目中，这两部分工作常常是耦合在一起的，由开发人员进行实现。
受限于经验/个人偏好，不同的开发人员进行技术选型、实现后，创造的代码会存在多层面的差异，包括：分层不统一、测试不友好、实现过程存在缺陷等。
即使同一个开发人员，TA在实现不同应用领域的代码时，也会有大量重复的框架代码进行重新编写，或者进行代码复制。
不排除有部分公用的函数，可以抽取成公共库被进行引用，但仍避免不了大量类似的领域逻辑代码进行重新编写。

如果使用代码生成器，可以专注在：1.领域知识建模；2. 框架模版创建。对于重复的框架代码，可以沉淀到框架模版中，实现一次编写，多次生成。

## 目前是否存在成熟解决方案？
据了解，[Telosys](http://www.telosys.org/)具备了上述代码生成器的能力。其基本框架如下图所示：
![](https://modeling-languages.com/wp-content/uploads/2013/07/Telosys-Overview.png)

## 为什么有 dddgen 这个代码库？
在最开始进行代码编写时，我个人的思路是将领域建模与框架代码分离开。
建模过程使用更高级的领域实体、关系进行描述，通过领域实体描述信息，自动生成符合 DDD 分层模型（符合工程实践）的业务代码。
因此着手实现了自己的小工具，并放到 github 之上。

## 本代码库后续如何运作？
代码生成一定会是效率提升的手段（无论是使用 Telosys 或是自己造的轮子），因此本代码库仍会更新。
在试用[apache velocity](https://velocity.apache.org/engine/1.7/user-guide.html#velocity-template-language-vtl-an-introduction)，
并对比 [go text/template](https://pkg.go.dev/text/template) 的基础上，对框架模版使用的技术进行再次选型。
在对比 [telosys model](https://doc.telosys.org/dsl-model) 和 [golang ast tag](https://pkg.go.dev/go/ast) 的基础上，对建模语言/方式进行再次选型。

