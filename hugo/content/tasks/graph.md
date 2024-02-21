---
menu:
    after:
        name: graph
        weight: 1
title: Построение графа
---

# Построение графа

Нужно написать воркер, который будет строить граф на текущей странице, каждые 5 секунд
От 5 до 30 элементов, случайным образом. Все ноды графа должны быть связаны.

```go
type Node struct {
 ID int
 Name string
	Form string // "circle", "rect", "square", "ellipse", "round-rect", "rhombus"
 Links []*Node
}
```

## Mermaid Chart

[MermaidJS](https://mermaid-js.github.io/) is library for generating svg charts and diagrams from text.

## Пример

{{< columns >}}
```tpl
{{</*/* mermaid [class="text-center"]*/*/>}}
graph LR
Main-->Node0
Node0-->Node1
Node0-->Node2
Node0-->Node3
Node1-->Node2
Node1-->Node3
Node2-->Node3

{{</*/* /mermaid */*/>}}
```

<--->

{{< mermaid >}}
graph LR
Main-->Node0
Node0-->Node1
Node0-->Node2
Node0-->Node3
Node1-->Node2
Node1-->Node3
Node2-->Node3

{{< /mermaid >}}

{{< /columns >}}
