import java.math.BigInteger;
import java.util.UUID;

import com.github.f4b6a3.uuid.UuidCreator;

public class App {
    public static void main(String[] args) {
        // Generate a UUID version 7
        UUID uuid = UuidCreator.getTimeOrderedEpoch();

        System.out.println("uuid: " + uuid);

        // Convert the UUID to a BigInteger with base 16
        BigInteger tokenId = new BigInteger(UuidCreator.toString(uuid).replace("-", ""), 16);

        System.out.println("tokenId: " + tokenId.toString());
    }
}