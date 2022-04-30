# DungeonFinder 

Dungeon finder is both a fancy technical Hello-World combined with the attempt to create a Full Stack application to manage real-life DnD Groups and Sessions. 

The use case comes from the lack of a german real-life group manager combined with some technical playground to create a few microservices to learn and practise language that I / we dont use professionally such as Rust, Golang or Svelte. Moreover we will also play around with different technologies and figure out where advantages and disadvantages are first hand instead of always trusting the documentation. 


```plantuml 
@startmindmap
* **Dungeon Finder**
** UI - Svelte
*** Group
*** Session 
*** User Profile 
**** DM Profile 
**** Player Profile 

** Backend - Micro services
***  User Services
**** RealUser Service
*****:DB Model
<code>
id uuid 
email string
fist-name string 
last-name string
password-salt string 
password hash
created datetime 
updated datetime 
dm-id uuid 
player-id uuid 
</code>
;
**** Player Service
*****:DB Models

**PlayerProfile**
<code>
id uuid 
characters: array{character}
</code>

**Character**
<code>
id uuid 
name string 
description string  
stats  object{stats}
</code>

**Stats** 
<code>
strength object{Attribute}
constitution object{Attribute}
initiative object{Attribute}
wisdom object{Attribute}
intelligence object{Attribute}
charisma object{Attribute}
</code>

**Attribute** 
<code>
Value uint 
Modifier uint 
</code>
;
**** DM Service
*** Session Service 
*** Group Service 

@endmindmap
```