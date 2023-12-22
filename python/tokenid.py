import uuid6

if __name__ == "__main__":
    uuid_id = uuid6.uuid7()
    print(f"UUID: {uuid_id}")

    # Convert the UUID to uint256
    token_id = int(uuid_id.hex, 16)
    print(f"Token ID: {token_id}")