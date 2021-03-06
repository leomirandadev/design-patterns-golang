# Design Patterns GoF
Padrões de projeto implementados em Golang. Estão abordados os padrões:
- Criacionais (creational);
- Estruturais (structural);
- Comportamentais (behavioral);


# Padrões Criacionais (creational)

## ⭐ Abstract Factory

Fornece uma interface para criação de famílias de objetos relacionados ou dependentes sem especificar suas classes concretas.

## ⭐ Factory Method

Define uma interface para criar um objeto, mas deixa as subclasses decidirem qual classe a ser instanciada. O Factory Method permite a uma classe postergar (defer) a instanciação às classes.

## ⭐ Singleton

Garante que uma classe tenha somente uma instância e fornece um ponto global de acesso para ela.

# Padrões Estruturais (structural)

## ⭐ Adapter

Converte a interface de uma classe em outra interface esperada pelos clientes. O adapter permite que certas classes trabalhem em conjunto, pois de outra forma seria impossível por causa de suas interfaces incompatíveis.

## ⭐ Composite

Compões objetos em estrutura de árvore para representar hierarquias do tipo partes-todo. O composite permite que os clientes tratem objetos individuais e composições de objetos de maneira uniforme.

## ⭐ Decorator

Atribui responsabilidades adicionais a um objeto dinamicamente. Os decorators fornecem uma alternativa flexível a subclasses para extensão da funcionalidade.

## ⭐ Façade

Fornece uma interface unificada para um conjunto de interfaces em um subsistema. O Façade define uma interface de nível mais alto que torna o subsistema mais fácil de usar.

# Padrões Comportamentais (behavioral)

## ⭐ Observer

Define uma dependência um-para-muitos entre objetos, de modo que, quando um objeto muda de estado, todos os seus dependentes são automaticamente notificados e atualizados.

## ⭐ Strategy

Define uma família de algoritmos, encapsular cada um deles e fazê-los intercambiáveis. O Strategy permite que o algoritmo varie independentemente dos clientes que o utilizam

## ⭐ Template Method

Define o esqueleto de um algoritmo em uma operação, postergando a definição de alguns passos para subclasses. O Template Method permite que as subclasses redefinam certos passos de um algoritmo sem mudar sua estrutura.