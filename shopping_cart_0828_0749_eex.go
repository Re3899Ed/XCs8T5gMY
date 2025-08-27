// 代码生成时间: 2025-08-28 07:49:46
package controllers

import (
    "encoding/json"
    "github.com/revel/revel"
)

// ShoppingCartController handles shopping cart functionality.
type ShoppingCartController struct {
    *revel.Controller
}

// AddItem adds a new item to the shopping cart.
func (c ShoppingCartController) AddItem(
    itemId string, quantity int,
) revel.Result {
    // Retrieve the shopping cart from the session.
    cart, ok := c.Session["cart"]
    if !ok {
        // Initialize a new shopping cart if not exists.
        cart = make(map[string]int)
    }

    // Add or update the item quantity in the cart.
    if cart != nil {
        cart.(map[string]int)[itemId] += quantity
    }

    // Save the updated cart back to the session.
    c.Session["cart"] = cart

    // Return a success message.
    return c.RenderJSON(
        map[string]string{"message": "Item added successfully"},
    )
}

// RemoveItem removes an item from the shopping cart.
func (c ShoppingCartController) RemoveItem(itemId string) revel.Result {
    // Retrieve the shopping cart from the session.
    cart, ok := c.Session["cart"]
    if !ok {
        // Return an error message if the cart doesn't exist.
        return c.RenderJSON(
            map[string]string{"error": "Cart not found"},
        )
    }

    // Remove the item from the cart.
    if cart != nil {
        delete(cart.(map[string]int), itemId)
    }

    // Save the updated cart back to the session.
    c.Session["cart"] = cart

    // Return a success message.
    return c.RenderJSON(
        map[string]string{"message": "Item removed successfully"},
    )
}

// GetCart retrieves the current shopping cart contents.
func (c ShoppingCartController) GetCart() revel.Result {
    // Retrieve the shopping cart from the session.
    cart, ok := c.Session["cart"]
    if !ok {
        // Return an empty cart if not exists.
        return c.RenderJSON(map[string]int{})
    }

    // Return the current cart contents.
    return c.RenderJSON(cart)
}
