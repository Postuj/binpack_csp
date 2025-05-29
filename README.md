# Binpack CSP

This is a student project for solving the Bin Packing Problem using Constraint Satisfaction Problem (CSP) techniques.

## Problem Description

This is a problem of packing items into proper bins. The goal is to pack all items into the bins while satisfying a set of constraints.
The problem is defined as follows:

Given a set of bins with various types and sizes:

| Name           | Type    | Capacity |
| -------------- | ------- | -------- |
| Crate          | regular | 16       |
| Box            | regular | 16       |
| Crate with ice | cooled  | 32       |
| Cooled box     | cooled  | 32       |
| Fridge         | cooled  | 32       |


and a set of items with different types, sizes and allowed bin types:

| Name    | Type      | Size | Allowed bin type |
| ------- | --------- | ---- | ---------------- |
| Apple   | fruit     | 4    | regular          |
| Banana  | fruit     | 8    | regular          |
| Orange  | fruit     | 4    | regular          |
| Potato  | vegetable | 8    | regular          |
| Onion   | vegetable | 4    | regular          |
| Tomato  | vegetable | 4    | regular          |
| Salmon  | seafood   | 16   | cooled           |
| Tuna    | seafood   | 24   | cooled           |
| Shrimp  | seafood   | 8    | cooled           |
| Crab    | seafood   | 8    | cooled           |
| Pork    | meat      | 16   | cooled           |
| Chicken | meat      | 8    | cooled           |
| Beef    | meat      | 8    | cooled           |


the goal is to place items in the bins in such a way that all of the following constraints are satisfied:

- Each item is allocated exactly once
- Each item is allocated in one bin only
- Each item is allocated on continuous slots in the bin
- Each item is allocated within the bin boundaries
- Items are allocated without overlapping
- Items of type `fruit` and `vegetable` can be placed in regular bins only
- Items of type `seafood` and `meat` can be placed in cooled bins only
- Items of type `fruit` and `vegetable` cannot be placed together in the same bin
- Items of type `seafood` and `meat` cannot be placed together in the same bin
- All items are placed in bins

## Running

### Requirements

1. Go version 1.20 or higher

### Steps

1. `go run ./app`
