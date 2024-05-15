package main

// Cons represents a non-empVty list
type Cons[T any] struct {
    headItem T
    tailList LList[T]
}

func (Cons[T]) isEmpty() bool {
    return false
}

func (c Cons[T]) head() (T, error) {
    return c.headItem, nil
}

func (c Cons[T]) tail() (LList[T], error) {
    return c.tailList, nil
}

func (c Cons[T]) length() int {
    return 1 + c.tailList.length()
}

func (c Cons[T]) copy() LList[T] {
    return Cons[T]{c.headItem, c.tailList.copy()}
}

func (c Cons[T]) append(list LList[T]) LList[T] {
    return Cons[T]{c.headItem, c.tailList.append(list)}
}

func (c Cons[T]) transform(f func(interface{}) interface{}) LList[T] {
    // Transform the head item using the provided function
    transformedHead := f(c.headItem)

    // Assert the transformed head item back to type T
    head, ok := transformedHead.(T)
    if !ok {
        // Handle the case where the type assertion fails
        // For example, you might return an error or panic
    }

    // Transform the tail list recursively
    transformedTail := c.tailList.transform(f)

    // Return the new Cons with the transformed head and tail
    return Cons[T]{head, transformedTail}
}