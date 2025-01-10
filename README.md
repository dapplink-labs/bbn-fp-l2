# Manta FP

Manta FP (Fast Finality Protocol) is a blockchain solution designed to enable fast finality and quick withdrawal mechanisms while ensuring network security. It leverages a dual-protocol staking model to optimize the finality of Layer-2 blocks and reduce fraud-proof times for Layer-1 withdrawals.

## Features

- **Fast Layer-2 Finality**: Accelerates block confirmation on the Manta Layer-2 network.
- **Quick Layer-1 Withdrawals**: Minimizes the time required for fraud proofs, enabling rapid token withdrawals.
- **Dual-Protocol Staking Model**:
    - *Babylon*: BTC staking-based finality provider.
    - *Symbiotic*: Ethereum-based re-staking protocol.
- **Secure and Efficient Architecture**:
    - Relayer network for cross-layer and cross-protocol coordination.
    - Slashing and reward mechanisms to ensure network reliability.

## Core Components

1. **Babylon Finality Provider**:
    - BTC stakers delegate vote power to Manta FP nodes.
    - FP nodes sign state roots and submit aggregated signatures to Babylon for validation.

2. **Symbiotic Finality Provider**:
    - Manta stakers delegate vote power to Manta FP nodes.
    - FP nodes submit aggregated state root signatures to the Manta Relayer.

3. **Manta Relayer Network**:
    - Aggregates and verifies signatures from both protocols.
    - Submits state roots and related signatures to Layer-1 contracts (e.g., Ethereum) for quick withdrawals.

4. **Slashing and Rewards**:
    - Incentivizes honest behavior and penalizes faulty nodes.
    - Rewards are distributed for successful validation, while misbehavior leads to slashing.

## Architecture

- **FP Nodes**: Handle state root commitments and signature aggregation.
- **Relayer Network**: Bridges Babylon, Symbiotic, and Layer-1 protocols.
- **Smart Contracts**:
    - `MantaManagerService`: Validates commitments and manages staking rewards.
    - `DGF` or equivalent contracts: Handles Layer-1 fraud-proof reductions.

## Getting Started

### Prerequisites

- Golang version 1.20 or higher.
- A properly configured environment for building and running blockchain applications.

### Installation

1. Clone the repository:
   ```bash
   git clone https://github.com/your-repo/manta-fp.git
   ```
2.	Navigate to the project directory:
   ```bash
  cd manta-fp
   ```
Usage

1. Start the Manta FP service:

   ```bash
    TODO
   ```

Contributing

We welcome contributions! Please follow these steps:
1.	Fork the repository.
2.	Create a feature branch: git checkout -b feature-name.
3.	Commit your changes: git commit -m "Add feature".
4.	Push to the branch: git push origin feature-name.
5.	Open a pull request.

License

This project is licensed under the MIT License. See the LICENSE file for details.

Contact

For questions or support, please reach out to us at support@manta.network or visit our [website](https://manta.network/).