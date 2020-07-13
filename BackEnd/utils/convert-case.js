import jsConvertCase from 'js-convert-case';

export function camelKeysArray(array) {
  var camelArray = [];
  for (let element of array) {
    camelArray.push(jsConvertCase.camelKeys(element));
  }
  return camelArray;
}
