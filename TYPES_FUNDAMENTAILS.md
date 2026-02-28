In programming, you would work with data, either by manipulating existing or just saving it or deleting, point is, you would always have data to work with.

Now, first thing is, the data can't just be floating in the air, it needs to be stored somewhere and how you store them is subject to the use case.
Do you need the data to be permanently stored during the lifecycle of the application with no modification or it would dynamically change as the application runs?
the difference here defines if the data would change or not- for this we use the term "mutable" and "immutable" data which for a program like Go can be stored as constants or variables.

Now, what are stored regardless of their durability or mutability, these are referred to as data types.

Go has a number of built-in data types, but you can also create your own data types - this is called custom data types.
Working with data types is dependent on the data or value you want "stored", if it a number value, you have between integers and floating point numbers, for texts, alphanumeric or not, you have strings. If it is a true or false value, you have booleans. If it is a collection of other values, you have arrays and slices, maps, and structs.
Lastly, a pointer is a data type which is used to point to a value and also, a value can be nil which is a null value.
All of the above can also still be used together to create more complex custom data types.

Next, we would talk about working with data types in Go - which is called functions. This can be custom functions as well or built-in functions which are referred to as methods specifically designed for working with data types that they are associated with.
