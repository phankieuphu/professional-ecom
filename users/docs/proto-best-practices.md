
Certainly! Here are some practical examples illustrating the best practices for Protocol Buffers (Protobuf) as outlined in the official documentation:

---

### ✅ Do Reserve Tag Numbers for Deleted Fields

**Bad Example:**

```proto
message User {
  string name = 1;
  int32 age = 2;
  string email = 3; // Removed in future versions
}
```

**Good Example:**

```proto
message User {
  string name = 1;
  int32 age = 2;
  reserved 3; // Reserved tag number to prevent reuse
}
```

*Reserving tag numbers ensures that deleted fields' tag numbers are not accidentally reused, preventing potential deserialization issues.*

---

### ✅ Do Include an Unspecified Value in an Enum

**Bad Example:**

```proto
enum Status {
  ACTIVE = 1;
  INACTIVE = 2;
}
```

**Good Example:**

```proto
enum Status {
  STATUS_UNSPECIFIED = 0; // Default value
  ACTIVE = 1;
  INACTIVE = 2;
}
```

*Including a default `UNSPECIFIED` value allows for future enum additions without breaking existing clients.*

---

### ✅ Do Use Well-Known Types and Common Types

**Bad Example:**

```proto
message Event {
  int64 timestamp = 1;
  int32 duration_seconds = 2;
}
```

**Good Example:**

```proto
import "google/protobuf/timestamp.proto";
import "google/protobuf/duration.proto";

message Event {
  google.protobuf.Timestamp timestamp = 1;
  google.protobuf.Duration duration = 2;
}
```

*Utilizing well-known types like `Timestamp` and `Duration` promotes consistency and interoperability.*

---

### ✅ Do Define Message Types in Separate Files

**Bad Example:**

```proto
// user.proto
message User {
  string name = 1;
  int32 age = 2;
  // Other fields...
}

// address.proto
message Address {
  string street = 1;
  string city = 2;
  // Other fields...
}
```

**Good Example:**

```proto
// user.proto
import "address.proto";

message User {
  string name = 1;
  int32 age = 2;
  Address address = 3;
}

// address.proto
message Address {
  string street = 1;
  string city = 2;
  // Other fields...
}
```

*Separating message types into distinct files enhances modularity and maintainability.*

---

### ❌ Don’t Re-use a Tag Number

**Bad Example:**

```proto
message User {
  string name = 1;
  int32 age = 2;
  string email = 3; // Removed in future versions
  string phone = 3; // Reusing tag number 3
}
```

*Reusing tag numbers can cause deserialization errors and data corruption.*

---

### ❌ Don’t Change the Type of a Field

**Bad Example:**

```proto
message User {
  string name = 1;
  int32 age = 2;
  bool is_active = 3; // Changed from int32 to bool
}
```

*Changing a field's type can lead to data loss or incorrect interpretations.*

---

For more detailed guidelines and additional best practices, you can refer to the official Protobuf documentation: ([Protocol Buffers][1]).

---

[1]: https://protobuf.dev/best-practices/dos-donts/?utm_source=chatgpt.com "Proto Best Practices | Protocol Buffers Documentation"
