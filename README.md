# take

This is boring:

```code go
ids := make([]int64, 0, len(items))

for _, item := range items {
	ids == append(ids, item.Id)
}
```

Even without generic, we could at least do:

```code go
ids := take.Int64(items, func(i int) int64 {
	return items[i].Id
})
```

There is also `take.String`, `take.Int`.
