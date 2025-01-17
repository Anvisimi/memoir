# Memoir Chain

A blockchain for storing and managing stories with CRUD (Create, Read, Update, Delete) operations.

## Overview
The chain allows users to:
- Create stories with title and content
- Read stories (list all or show specific)
- Update existing stories
- Delete stories

## Consensus Breaking Changes Demonstration

### What is a Consensus Breaking Change?
A consensus breaking change occurs when modifications to the blockchain code cause nodes running different versions to disagree on:
- State validation rules
- Transaction processing
- Block validation
- State transitions

Such changes lead to network splits (forks) as nodes cannot reach agreement on the chain's state.

### Our Consensus Breaking Changes
In the `feature/consensus-breaking-change` branch, we introduced:
1. New required fields in Story structure (`category` and `rating`)
2. Modified ID generation logic (incrementing by 2 instead of 1)

These changes break consensus because:
- Old nodes can't validate stories with new required fields
- New nodes reject old stories missing required fields
- Different ID sequences cause state divergence
- Existing state becomes invalid in new version

## How to Validate

### 1. Run Original Version
```bash
# Clone repository
git clone <repository-url>
cd memoir

# Switch to develop branch
git checkout develop

# Start the chain
ignite chain serve

# In a new terminal, create a story
memoird tx memoir create-story "Original Story" "Content" \
    --from bob \
    --chain-id memoir \
    --gas 300000 \
    --fees 3000stake \
    --keyring-backend test \
    --node tcp://localhost:26657 \
    -y

# Query stories - note the ID sequence and structure
memoird q memoir list-story --node tcp://localhost:26657
```

### 2. Run Modified Version
```bash
# Switch to consensus-breaking branch
git checkout feature/consensus-breaking-change

# Start the chain
ignite chain serve

# Create a new story
memoird tx memoir create-story "New Story" "Content" \
    --from bob \
    --chain-id memoir \
    --gas 300000 \
    --fees 3000stake \
    --keyring-backend test \
    --node tcp://localhost:26657 \
    -y

# Query stories - note different ID sequence and new fields
memoird q memoir list-story --node tcp://localhost:26657
```

### Observable Differences
1. **Story Structure**:
   - Original: Basic story with title, content, author
   - Modified: Additional required fields (category, rating)

2. **ID Generation**:
   - Original: Sequential (0, 1, 2...)
   - Modified: Increments by 2 (0, 2, 4...)

3. **State Compatibility**:
   - Old stories become invalid in new version
   - New stories can't be processed by old version

### Validation Points
An assessor can verify:
1. Different story structures between versions
2. Changed ID sequence
3. State incompatibility between versions
4. Network split potential due to validation rule changes

This demonstrates a valid consensus-breaking change as it modifies:
- Data structures
- State validation rules
- State transition logic
- Block processing rules

These changes would require a coordinated upgrade in a production environment.
