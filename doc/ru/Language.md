# Язык Anzer

Anzer является предметно-ориентированным языком, предназначеныным для описания взаимодействия
функций в FaaS платформах. Ближайшими его аналогами можно считать различные системы объединения
функций в последовательности: AWS State Machine, OpenWhisk Sequences и прочие.

Ключевое отличие языка - описание типов передаваемых данных и проверка типов на всех этапах
разработки систем для FaaS-платформ. Ближайшим аналогом языка среди языков программирования 
являются Haskell и PureScript, однако, Anzer - это не язык программирования общего пользования.
На языке Anzer нельзя написать какую-либо логику, язык предназначен для композиции функций,
написанных на других языках.

С помощью языка можно описать типы данных, передаваемых между функциями, сами функции и то, как
они взаимодействуют.