# The SchemaCAN Manifesto

Note: the idea for doing this was stolen from the [Redis Manifesto](https://github.com/redis/redis/blob/unstable/MANIFESTO).

## 1. Use tools for what they're made to do
Where possible, don't design outside of your tool. That is not to say don't try learn mastery of tools and only stick to a subset of features in a language or framework, but to make choice design decisions take are complemented by the toolset available to you.

## 2. The API needs to map well to the real thing
Lucky in our case we already have a strict interface that we must comply with and that is the CAN Bus specification (most notably ISO 11898). There is not much point in trying to abstract away this interface too much given the intended target audience is people who are close to the metal in their embedded systems anyway. Make your interface match reality and you will be rewarded.

## 3. Immitate tools with good ideas
With not much experience to draw from as this is a relatively new project, draw on the lessons learned by projects with more years under their belt. SchemaCAN is largely constructed around the ideals of Kubernetes manifests, which in turn are constructed around the Kubernetes API. This is no mistake and there are pleanty of other ideas stolen like the type names being consice and meaningful like those from the Zig programming language.

## 4. Operate on the basis of continual improvement, but be careful in your order of optomisation
It's not hard to convince oneself of why continual improvement is important, but it is easy to get caught up optomising systems that provide little overall benifit in the larger picture. When optomising, follow these steps _in order_ to reduce the amount of time wasted:
   1. Improve the requirements you must satisfy.
   2. Remove surplus components that no longer make sense in the current design.
   3. Simplify or optomise components or processes.
   4. Accelerate feeback loops.
   5. Automate.
