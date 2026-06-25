Exercises
---

## 1. Structs (uses: variables, functions, loops)

**Theme: Build a "Game Character Creator"**

### Exercise 1.1 — Character Stats
Create a `Character` struct with fields: `Name` (string), `Health` (int), `Strength` (int), `Level` (int). Write a function `NewCharacter(name string) Character` that returns a character with default stats (Health: 100, Strength: 10, Level: 1). Write another function `DisplayStats(c Character)` that prints a formatted character sheet.

### Exercise 1.2 — Level Up System
Add an `Experience` field (int). Write a method `GainXP(c *Character, amount int)` that adds XP. Every 100 XP, the character levels up (Level += 1, Strength += 2, Health += 20, and remaining XP carries over). Print what happened: *"Gained 150 XP! Level 2 reached! Strength +2, Health +20"*.

### Exercise 1.3 — Inventory with Slices (introduces slices early)
Add an `Inventory` field to `Character` as a `[]string`. Write methods `AddItem(c *Character, item string)` and `ShowInventory(c Character)`. Prevent duplicate items (if "Sword" is already there, print *"Already have Sword"*). Use a loop to check for duplicates.

### Exercise 1.4 — Party System
Create a `Party` struct with a `Members []Character` field. Write methods `AddMember(p *Party, c Character)`, `TotalPartyHealth(p Party) int`, and `FindStrongest(p Party) Character`. Use loops to iterate through the party.

---

## 2. Interfaces (uses: structs, methods, loops)

**Theme: "RPG Combat System — Different Attack Styles"**

### Exercise 2.1 — Define a Combatant Interface
Create an interface `Combatant` with methods `Attack() int` and `Defend() int`. Make `Character` (from structs) implement it: Attack returns Strength, Defend returns Health/10.

### Exercise 2.2 — New Enemy Type
Create a `Dragon` struct with `FirePower int` and `Scales int`. Implement `Combatant` for Dragon: Attack returns FirePower, Defend returns Scales. Write a function `BattleReport(c Combatant)` that prints: *"This combatant attacks for X and defends with Y"* — it should work with BOTH Character and Dragon.

### Exercise 2.3 — Battle Arena
Create a `BattleArena` that takes two `Combatant`s and simulates a fight. Each round, both attack each other (reduce some health value). The arena should accept ANY type that implements `Combatant` — a Character vs Dragon, two Dragons, etc. Run until one "dies" (health <= 0).

### Exercise 2.4 — Healer Interface (composition practice)
Create a `Healer` interface with `Heal() int`. Make a `Cleric` struct that implements BOTH `Combatant` and `Healer`. Write a function `SupportRound(c Combatant, h Healer)` where the healer heals the combatant before they attack. This teaches interface composition.

---

## 3. Errors (uses: structs, interfaces, loops)

**Theme: "A Bank with Real-World Problems"**

### Exercise 3.1 — Custom Error Types
Create a `BankAccount` struct with `Owner string` and `Balance float64`. Define custom error types: `InsufficientFundsError`, `InvalidAmountError`, and `AccountFrozenError` (all as structs implementing `error`). Each should have meaningful `Error()` messages.

### Exercise 3.2 — Withdraw with Error Handling
Write `Withdraw(account *BankAccount, amount float64) error`. Return `InvalidAmountError` if amount <= 0, `InsufficientFundsError` if amount > Balance. In `main`, attempt multiple withdrawals and handle each error type differently (print different messages based on error type).

### Exercise 3.3 — Error Wrapping
Create an `AccountManager` that manages multiple accounts. Write `Transfer(from, to *BankAccount, amount float64) error`. If `Withdraw` fails, wrap that error with `fmt.Errorf("transfer failed: %w", err)`. In `main`, use `errors.Is()` and `errors.As()` to unwrap and inspect the original error.

### Exercise 3.4 — Validation Pipeline
Write a function `ValidateTransaction(account BankAccount, amount float64, isFrozen bool) []error`. Return ALL validation errors found (not just the first one): negative amount, insufficient funds, frozen account. Collect them in a slice and return. The caller decides how to handle multiple errors.

---

## 4. Loops (uses: variables, functions, structs)

**Theme: "Pattern Generation & Simulation"**

### Exercise 4.1 — ASCII Art Generator
Write a function `DrawTriangle(height int)` that prints a right-aligned triangle using `*`. Then `DrawDiamond(height int)` that prints a diamond. Use nested loops. No recursion allowed — pure loop practice.

### Exercise 4.2 — Number Guessing Game
Implement a guessing game where the computer picks a random number (1-100). The player gets 7 guesses. After each guess, print "too high" or "too low". Use a `for` loop with a condition. Track all guesses in a slice and show them at the end.

### Exercise 4.3 — Conway's Game of Life (Simplified)
Create a 10x10 grid (2D slice of bools). Initialize with a "glider" pattern. Write one generation update function using nested loops. Print the grid using `*` for alive, `.` for dead. Run 5 generations and print each. This is a classic loop exercise that feels like a real simulation.

### Exercise 4.4 — Text Spinner Animation
Write a function that prints a spinning character (`|`, `/`, `-`, `\`) in place on the terminal. Use a loop with `fmt.Printf("\r%c", spinner[i%4])` and `time.Sleep(100 * time.Millisecond)`. Run for 3 seconds. This teaches loop control and terminal manipulation.

---

## 5. Slices (uses: loops, structs, functions)

**Theme: "A Music Playlist Manager"**

### Exercise 5.1 — Dynamic Playlist
Create a `Song` struct with `Title`, `Artist`, `Duration` (seconds). Create a `Playlist` type as `[]Song`. Write functions: `AddSong`, `RemoveSong(by index)`, `Shuffle` (using `rand.Shuffle`). Implement `TotalDuration(p Playlist) int` using a loop.

### Exercise 5.2 — Smart Filtering
Write `FilterByArtist(p Playlist, artist string) Playlist` that returns a NEW slice with matching songs (don't modify original). Write `FilterLongSongs(p Playlist, minDuration int) Playlist`. Practice slice creation and `append`.

### Exercise 5.3 — Playlist Merge & Deduplication
Write `MergePlaylists(p1, p2 Playlist) Playlist` that combines two playlists, removing duplicate songs (same Title + Artist). Use a map (if you've covered it) or nested loops (if not yet). Return a new slice.

### Exercise 5.4 — Slice Tricks — Undo System
Implement an undo feature for your playlist. Maintain a `history []Playlist` (slice of slices). Every time you modify the playlist, append a copy to history. `Undo()` removes the last state and restores it. This teaches slice copying (`append([]Song(nil), p...)`).

---

## 6. Maps (uses: slices, loops, structs)

**Theme: "A Restaurant Order System"**

### Exercise 6.1 — Menu & Prices
Create a `Menu` type: `map[string]float64` (item -> price). Write functions: `AddItem`, `GetPrice`, `HasItem`. Handle the "zero value" problem — distinguish between "price is $0" and "item doesn't exist".

### Exercise 6.2 — Order Tracking
Create an `Order` struct with `Items map[string]int` (item name -> quantity). Write `AddToOrder(o *Order, item string, qty int)`, `RemoveFromOrder`, and `CalculateTotal(order Order, menu Menu) float64`. Handle errors for invalid items.

### Exercise 6.3 — Frequency Analysis
Write a function `WordFrequency(text string) map[string]int` that counts how many times each word appears. Ignore case and punctuation. Return the map sorted by frequency (you'll need to extract to a slice and sort it). This is a classic map exercise with a twist.

### Exercise 6.4 — Inverted Index (Mini Search Engine)
Build an inverted index: `map[string][]int` where keys are words and values are slice of document IDs they appear in. Create 5 "documents" (strings). Write `Search(index map[string][]int, word string) []int` that returns all document IDs containing that word. Handle multiple word searches with AND logic (intersection of slices).

---

## 7. Advanced Functions (uses: all above)

**Theme: "Functional Programming in Go"**

### Exercise 7.1 — Higher-Order Functions
Write `MapSlice[T any](slice []T, transform func(T) T) []T` (before generics, do it for `[]int` and `[]string` separately). Then write `FilterSlice[T any](slice []T, predicate func(T) bool) []T`. Test by doubling all numbers in a slice, then filtering for even numbers.

### Exercise 7.2 — Function Composition
Write `Compose(f, g func(int) int) func(int) int` that returns a new function `f(g(x))`. Create three simple functions: `Add2`, `Multiply3`, `Square`. Compose them: `Compose(Compose(Add2, Multiply3), Square)` and verify the math.

### Exercise 7.3 — Closures — Counter Factory
Write `MakeCounter(start int) func() int` that returns a closure. Each call increments and returns the next value. Create two independent counters and show they don't interfere. Then make `MakeStepCounter(start, step int) func() int`.

### Exercise 7.4 — Callback System — Event Handler
Create a `Button` struct with `OnClick []func()`. Write `RegisterCallback(b *Button, fn func())` and `Trigger(b Button)`. When triggered, all registered functions run. Build a simple GUI simulation: "Login Button" that validates input, logs the attempt, and shows a welcome message — all as separate callbacks.

---

## 8. Pointers (uses: all above)

**Theme: "In-Place Modifications & Memory Efficiency"**

### Exercise 8.1 — Swap Without Return
Write `Swap(a, b *int)` that swaps two integers in place. No return values. Test with `x, y := 5, 10`. Then write `SwapStrings(a, b *string)`.

### Exercise 8.2 — In-Place Slice Reversal
Write `ReverseInPlace(s []int)` that reverses a slice WITHOUT creating a new slice. Use two pointers (indices) moving toward the center. Verify the original slice is modified.

### Exercise 8.3 — Pointer Receiver Methods
Refactor your `BankAccount` from the Errors section. Make all methods use pointer receivers: `Deposit`, `Withdraw`, `Transfer`. Create a `[]*BankAccount` slice. Write a function `ApplyInterest(accounts []*BankAccount, rate float64)` that adds interest to ALL accounts in the slice using pointers.

### Exercise 8.4 — Linked List (Manual)
Implement a singly linked list WITHOUT using a slice. Use a `Node` struct with `Value int` and `Next *Node`. Write `Append(head **Node, value int)` (pointer to pointer), `PrintList(head *Node)`, and `ReverseList(head **Node)` that reverses in place by manipulating pointers. This is the ultimate pointer exercise.

---

## 9. Channels & Concurrency (uses: all above)

**Theme: "A Concurrent Pipeline"**

### Exercise 9.1 — Number Pipeline
Create three goroutines connected by channels:
- **Generator**: sends numbers 1-100 to a channel
- **Doubler**: reads from generator, doubles the number, sends to next channel
- **Printer**: reads from doubler and prints

Use `make(chan int)` and close channels properly.

### Exercise 9.2 — Concurrent File "Downloader" (Simulated)
Create a `Download(url string) string` function that simulates downloading by sleeping for random time (100-500ms) and returning `"Downloaded: " + url`. Write `DownloadAll(urls []string) []string` that downloads all URLs concurrently using goroutines, collects results through a channel, and returns them. Ensure you don't miss any results.

### Exercise 9.3 — Worker Pool
Implement a worker pool with 3 workers. Jobs are integers (1-20). Each worker sleeps for `job * 100ms` then prints `"Worker X processed job Y"`. Use a `jobs` channel and a `results` channel. The main goroutine sends 20 jobs and collects 20 results. Use `sync.WaitGroup` to wait for workers to finish.

### Exercise 9.4 — Fan-Out / Fan-In
Generate 100 random numbers. Fan-out to 4 "squarer" goroutines (each reads from the same input channel). Fan-in their results back to a single channel. Print all squared results. Use `select` in the fan-in goroutine to merge multiple channels into one.

---

## 10. Mutexes (uses: channels, goroutines, pointers)

**Theme: "Shared State Done Right"**

### Exercise 10.1 — Concurrent Counter
Create a `SafeCounter` struct with `count int` and `sync.Mutex`. Implement `Increment()`, `Decrement()`, and `Get() int`. Launch 1000 goroutines that each increment 100 times. Verify the final count is exactly 100,000. Run WITHOUT the mutex to see the race condition.

### Exercise 10.2 — Concurrent Cache
Build a `Cache` struct with `data map[string]string` and a `sync.RWMutex`. Implement `Get(key string) (string, bool)` using `RLock()`, and `Set(key, value string)` using `Lock()`. Launch goroutines that read and write concurrently. Print cache stats.

### Exercise 10.3 — Bank Transfer with Mutexes
Revisit your `BankAccount`. Add a `sync.Mutex` field. Make `Transfer` thread-safe: lock both accounts (always in the same order to avoid deadlock!). Launch 10 goroutines that randomly transfer between 3 accounts. Ensure no money is lost or created.

### Exercise 10.4 — Rate Limiter (Token Bucket)
Implement a token bucket rate limiter. `RateLimiter` struct has `tokens int`, `maxTokens int`, and a mutex. `Allow() bool` checks if a token is available (decrements if yes). A background goroutine adds 1 token every 100ms up to `maxTokens`. Test by firing 100 requests — only ~10 should succeed immediately, others should wait or fail.

---

## 11. Generics (uses: all above)

**Theme: "Write Once, Use Everywhere"**

### Exercise 11.1 — Generic Stack (from scratch)
Implement `Stack[T any]` with `Push`, `Pop() (T, bool)`, and `Peek() (T, bool)`. Test with `Stack[int]`, `Stack[string]`, and `Stack[BankAccount]`. Ensure `Pop` on empty stack returns zero value and false.

### Exercise 11.2 — Generic Map Functions
Write `MapValues[K comparable, V any](m map[K]V) []V` that extracts all values. Write `InvertMap[K comparable, V comparable](m map[K]V) (map[V]K, error)` that swaps keys and values (error if duplicate values). Test with `map[string]int` and `map[int]string`.

### Exercise 11.3 — Generic Filter & Reduce
Write `Filter[T any](slice []T, predicate func(T) bool) []T` and `Reduce[T any, U any](slice []T, initial U, reducer func(U, T) U) U`. Use them to: filter even numbers from `[]int`, then reduce to sum. Filter songs longer than 3 minutes from `[]Song`, then reduce to total duration.

### Exercise 11.4 — Generic Concurrent Safe Map
Combine everything! Build `SafeMap[K comparable, V any]` using generics AND `sync.RWMutex`. Methods: `Set`, `Get`, `Delete`, `Len() int`, `Keys() []K`, `Values() []V`. Test with concurrent goroutines reading and writing `SafeMap[string, int]` and `SafeMap[int, Song]`. This is your capstone project.

---

## How to Use These

1. **Don't look at solutions first.** Spend 20-30 minutes struggling before checking hints.
2. **Build incrementally.** Each exercise in a section builds on the previous.
3. **Write tests.** Even simple `if got != want { t.Errorf(...) }` tests help solidify understanding.
4. **Mix and match.** After section 5, try combining exercises (e.g., use your Playlist with your Bank system for a "Music Store").
5. **Refactor as you learn.** When you hit interfaces, go back and make your structs implement useful interfaces.

These exercises are designed to feel like building real tools while forcing you to use Go's unique features. Good luck escaping tutorial hell! 🚀