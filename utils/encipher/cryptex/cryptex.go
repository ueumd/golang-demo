/*
 * 用来加密配置
 * 配置中敏感参数使用本程序加密后写进配置
 * 程序使用硬编码的密钥解密后使用
 *
 * wencan
 * 2018-08-31
 */

package cryptex

import "myapiserver/utils/encipher"

// key 硬编码在代码中的对称密钥，用来加密配置参数
var key = []byte("dFR0kah3zgfd2dsksNay7DhdCk")
// Cryptex 用来加解密配置
var Cryptex encipher.Cipher

func init() {
	if Cryptex == nil {
		Cryptex = encipher.NewAesEncipher(key)
	}
}

// Reset 重设加解密器
func Reset(_key []byte) {
	key = _key
	Cryptex = encipher.NewAesEncipher(_key)
}
