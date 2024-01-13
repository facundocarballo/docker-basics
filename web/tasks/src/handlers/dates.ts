export const ConvertUint8ToDateString = (bytes: Uint8Array): string => {
  const textDecoder = new TextDecoder("utf-8");
  return textDecoder.decode(bytes);
};
