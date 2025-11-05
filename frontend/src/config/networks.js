// Правильные ABI из JSON файлов
import bridgeAbiData from './bridge-abi.json' assert { type: 'json' };
import tokenAbiData from './token-abi.json' assert { type: 'json' };

export const CONFIG = {
  ETHEREUM: {
    chainId: '0xaa36a7',
    rpcUrl: 'https://sepolia.infura.io/v3/afa8f63b32a84f508f554b798b247453',
    bridgeAddress: '0xc2766cBFc1Dc3E95058bf09c929E7D6226b10187',
    tokenAddress: '0x5CFdE9C777be47FC4a401c918181DD92BA4c81Cc',
    tokenAbi: tokenAbiData,
    bridgeAbi: bridgeAbiData
  },
  POLYGON: {
    chainId: '0x13882',  
    rpcUrl: 'https://polygon-amoy.infura.io/v3/afa8f63b32a84f508f554b798b247453',
    tokenAddress: '0x5CFdE9C777be47FC4a401c918181DD92BA4c81Cc',
    tokenAbi: tokenAbiData
  }
};
