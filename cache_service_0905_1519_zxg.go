// 代码生成时间: 2025-09-05 15:19:51
maintainability and extensibility.
*/

package cache

import (
    "encoding/json"
    "errors"
    "fmt"
    "time"
)

// CacheService represents the caching service.
type CacheService struct {
    // Store holds the cache entries.
    Store map[string]*CacheEntry
    // TTL defines the time-to-live for cache entries.
    TTL time.Duration
}

// CacheEntry represents a cache entry with a value and expiration time.
type CacheEntry struct {
    Value     []byte
    ExpiresAt time.Time
}

// NewCacheService creates a new CacheService instance with the given TTL.
func NewCacheService(ttl time.Duration) *CacheService {
    return &CacheService{
        Store:   make(map[string]*CacheEntry),
        TTL:     ttl,
    }
}

// Set adds or updates a cache entry with a given key and value.
func (cs *CacheService) Set(key string, value interface{}) error {
    // Convert the value to JSON.
    data, err := json.Marshal(value)
    if err != nil {
        return err
    }
    // Calculate the expiration time.
    expiresAt := time.Now().Add(cs.TTL)
    // Add or update the cache entry.
    cs.Store[key] = &CacheEntry{Value: data, ExpiresAt: expiresAt}
    return nil
}

// Get retrieves a cache entry by its key.
func (cs *CacheService) Get(key string) (interface{}, error) {
    entry, exists := cs.Store[key]
    if !exists {
        return nil, errors.New("cache entry not found")
    }
    if time.Now().After(entry.ExpiresAt) {
        delete(cs.Store, key) // Remove expired entry.
        return nil, errors.New("cache entry expired")
    }
    // Unmarshal the JSON value.
    var value interface{}
    if err := json.Unmarshal(entry.Value, &value); err != nil {
        return nil, err
    }
    return value, nil
}

// Flush removes all cache entries.
func (cs *CacheService) Flush() {
    cs.Store = make(map[string]*CacheEntry)
}
