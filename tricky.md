  # Tricky Lock
  ```
  void Lock() {
    while (thread_count_.fetch_add(1) > 0) {
      thread_count_.fetch_sub(1);
    }
  }

  void Unlock() {
    thread_count_.fetch_sub(1);
  }
  ```
