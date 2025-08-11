// 代码生成时间: 2025-08-11 18:52:12
package shopping

import (
    "encoding/json"
    "fmt"
    "revel"
    "strings"
)

// Item represents a shopping cart item
type Item struct {
    ID          string `json:"id"`
    Name        string `json:"name"`
    Description string `json:"description"`
    Price       float64 `json:"price"`
}

// Cart represents a shopping cart
type Cart struct {
    Items       map[string]*Item `json:"items"`
    TotalPrice float64          `json:"totalPrice"`
}

// AddItem adds a new item to the cart
func (c *Cart) AddItem(item *Item) error {
    if item == nil {
        return fmt.Errorf("cannot add nil item to cart")
    }
    c.Items[item.ID] = item
    c.TotalPrice += item.Price
    return nil
}

// RemoveItem removes an item from the cart by ID
func (c *Cart) RemoveItem(itemID string) error {
    if _, exists := c.Items[itemID]; !exists {
        return fmt.Errorf("item with id %s does not exist in cart", itemID)
    }
    delete(c.Items, itemID)
    c.recalculateTotalPrice()
    return nil
}

// recalculateTotalPrice recalculates the total price of the cart
func (c *Cart) recalculateTotalPrice() {
    var total float64
    for _, item := range c.Items {
        total += item.Price
    }
    c.TotalPrice = total
}

// Controller for handling shopping cart requests
type CartController struct {
    revel.Controller
}

// Index action lists all items in the cart
func (c *CartController) Index() revel.Result {
    cart := Cart{
        Items:       make(map[string]*Item),
        TotalPrice: 0,
    }
    // Here you would normally fetch the cart items from a database
    // For simplicity, we're just adding a sample item
    item := &Item{
        ID:          "1",
        Name:        "Sample Item",
        Description: "This is a sample item",
        Price:       10.99,
    }
    if err := cart.AddItem(item); err != nil {
        c.RenderError(500, err)
        return nil
    }

    // Convert cart to JSON
    cartJSON, err := json.Marshal(cart)
    if err != nil {
        c.RenderError(500, err)
        return nil
    }

    return c.RenderJSON(cartJSON)
}

// Add action adds a new item to the cart
func (c *CartController) Add(itemName, itemDescription string, itemPrice float64) revel.Result {
    item := &Item{
        ID:          strings.NewUUID().String(),
        Name:        itemName,
        Description: itemDescription,
        Price:       itemPrice,
    }
    cart := Cart{
        Items: make(map[string]*Item),
    }
    // Normally, you would retrieve the existing cart from a database
    if err := cart.AddItem(item); err != nil {
        c.RenderError(500, err)
        return nil
    }

    // Convert cart to JSON
    cartJSON, err := json.Marshal(cart)
    if err != nil {
        c.RenderError(500, err)
        return nil
    }

    return c.RenderJSON(cartJSON)
}

// Remove action removes an item from the cart
func (c *CartController) Remove(itemID string) revel.Result {
    cart := Cart{
        Items: make(map[string]*Item),
    }
    // Normally, you would retrieve the existing cart from a database
    if err := cart.RemoveItem(itemID); err != nil {
        c.RenderError(500, err)
        return nil
    }

    // Convert cart to JSON
    cartJSON, err := json.Marshal(cart)
    if err != nil {
        c.RenderError(500, err)
        return nil
    }

    return c.RenderJSON(cartJSON)
}