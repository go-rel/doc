# Upgrading

## Upgrading from 0.16.x to 0.17.x

### UpdateAll and DeleteAll

`UpdateAll` and `DeleteAll` now renamed to UpdateAny and DeleteAny respectively.

=== "0.16.x"

    ```go
    // UpdateAll
    err := repo.UpdateAll(ctx, query, Set("notes", "notes"))

    // DeleteAll
    err := repo.DeleteAll(ctx, query)
    ```

=== "0.17.x"

    ```go
    // UpdateAny
    updatedCount, err := repo.UpdateAny(ctx, query, Set("notes", "notes"))

    // DeleteAny
    deletedCount, err := repo.DeleteAny(ctx, query)
    ```

## Upgrading from 0.13.x to 0.14.x

### UpdateAll and DeleteAll

`UpdateAll` and `DeleteAll` function now returns affected rows instead of just error.

=== "0.13.x"

    ```go
    // UpdateAll
    err := repo.UpdateAll(ctx, query, Set("notes", "notes"))
    
    // DeleteAll
    err := repo.DeleteAll(ctx, query)
    ```

=== "0.14.x"

    ```go
    // UpdateAll
    updatedCount, err := repo.UpdateAll(ctx, query, Set("notes", "notes"))
    
    // DeleteAll
    deletedCount, err := repo.DeleteAll(ctx, query)
    ```

## Upgrading from 0.8.x to 0.9.x

### Association Definition

Before 0.9.x all association is saved as is when its parent is modified, after this PR (https://github.com/go-rel/rel/pull/127) this feature needs to be explicitly enabled.

=== "0.8.x"

    ```go
    // User schema.
    type User struct {
        ID      int
        Name    string
        Address Address
    }
    ```

=== "0.9.x"

    ```go
    // User schema.
    type User struct {
        ID      int
        Name    string
        Address Address `autosave:"true"`
    }
    ```

## Upgrading from 0.7.x to 0.8.x

### Import path

REL is migrated to a new github organization and all import path need to be moved from `Fs02` to `go-rel`

=== "0.7.x"

    ```go
    import github.com/Fs02/rel
    ```

=== "0.8.x"

    ```go
    import github.com/go-rel/rel
    ```
