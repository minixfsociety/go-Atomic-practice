***Old Style Atomics: Summary Guide***
What is "Old Style"? This is the classic way of writing atomic operations in Go. It has been the industry standard since the language was released. It uses direct functions from the sync/atomic package (like atomic.AddInt64).

When was it popular? It was the only standard from 2009 to 2022. In Go 1.19 (August 2022), "typed" atomics (like atomic.Int64) were introduced as the modern style. However, millions of lines of code at companies like Google, Uber, and Cloudflare are still written in this "Old Style." Knowing this is a requirement for any professional Go engineer.

***Quick Command Reference***
atomic.AddInt64(&val, n) — "The Incrementer"
What it does: Safely adds number n to the variable val.
**Folder: atomic.add**
atomic.LoadInt64(&val) — "The Safe Reader"
What it does: Guarantees you read the absolute latest value directly from memory, bypassing any "stale" CPU caches.
**Folder: atomic.load**
atomic.StoreInt64(&val, n) — "The Switch"
What it does: Instantly overwrites the value. Used for global flags (e.g., "Shut down server now").
**Folder: atomic.store**
atomic.CompareAndSwapInt64(&val, old, new) — "The Bouncer"
What it does: Updates the value to new only if the current value is exactly old. The foundation of Leader Election.
**Folder: atomic.cas**
💡 Why do we use & everywhere?
Because in "Old Style," these functions need the memory address (pointer) of the variable so they can modify the original data directly in the RAM, rather than working on a local copy.
