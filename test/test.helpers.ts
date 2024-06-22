import { HanaTokenConsumer__factory } from "@typechain-types";
import { BigNumber, ContractReceipt } from "ethers";

export const parseHanaConsumerLog = (logs: ContractReceipt["logs"]) => {
  const iface = HanaTokenConsumer__factory.createInterface();

  const eventNames = logs.map((log) => {
    try {
      const parsedLog = iface.parseLog(log);

      return parsedLog.name;
    } catch (e) {
      return "NO_ZETA_LOG";
    }
  });

  return eventNames;
};
