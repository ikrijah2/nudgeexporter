/*
 * Licensed to the Apache Software Foundation (ASF) under one
 * or more contributor license agreements. See the NOTICE file
 * distributed with this work for additional information
 * regarding copyright ownership. The ASF licenses this file
 * to you under the Apache License, Version 2.0 (the
 * "License"); you may not use this file except in compliance
 * with the License. You may obtain a copy of the License at
 *
 *   http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing,
 * software distributed under the License is distributed on an
 * "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
 * KIND, either express or implied. See the License for the
 * specific language governing permissions and limitations
 * under the License.
 */
 /* jshint -W100 */

import { ThriftTest } from "./gen-js/ThriftTest_types";
import "./gen-js/ThriftTest";
var Int64 = require("node-int64");
var phantom = require("phantom");

const int64_2_pow_60: typeof Int64 = new Int64('1000000000000000');
const int64_minus_2_pow_60: typeof Int64 = new Int64('f000000000000000');

 (function() {
    'use strict';
  
    // Rudimentary test helper functions
    // TODO: Return error code based on kind of errors rather than throw
    var ok = function(t, msg) {
      if (!t) {
        console.log('*** FAILED ***');
        throw new Error(msg);
      }
    };
    var equal = function(a, b) {
      if (a !== b) {
        console.log('*** FAILED ***');
        throw new Error();
      }
    };
    var test = function(name, f) {
      console.log('TEST : ' + name);
      f();
      console.log('OK\n');
    };
  
    var parseArgs = function(args) {
      var skips = [
        '--transport=http',
        '--protocol=json'
      ];
      var opts = {
        port: '9090'
        // protocol: 'json',
      };
      var keys = {};
      for (var key in opts) {
        keys['--' + key + '='] = key;
      }
      for (var i in args) {
        var arg = args[i];
        if (skips.indexOf(arg) != -1) {
          continue;
        }
        var hit = false;
        for (var k in keys) {
          if (arg.slice(0, k.length) === k) {
            opts[keys[k]] = arg.slice(k.length);
            hit = true;
            break;
          }
        }
        if (!hit) {
          throw new Error('Unknown argument: ' + arg);
        }
      }
      var portAsInt: number = parseInt(opts.port, 10);
      if (!opts.port || portAsInt < 1 || portAsInt > 65535) {
        throw new Error('Invalid port number');
      }
      return opts;
    };
  
    var execute = function() {
      console.log('### Apache Thrift Javascript standalone test client');
      console.log('------------------------------------------------------------');
  
      phantom.page.injectJs('thrift.js');
      phantom.page.injectJs('gen-js/ThriftTest_types.js');
      phantom.page.injectJs('gen-js/ThriftTest.js');
  
      var system = require('system');
      var opts = parseArgs(system.args.slice(1));
      var port = opts.port;
      var transport = new Thrift.Transport('http://localhost:' + port + '/service');
      var protocol = new Thrift.Protocol(transport);
      var client = new ThriftTest.ThriftTestClient(protocol);
  
  
      // TODO: Remove duplicate code with test.js.
      // all Languages in UTF-8
      var stringTest = "Afrikaans, Alemannisch, Aragon??s, ??????????????, ????????, Asturianu, Aymar aru, Az??rbaycan, ??????????????, Boarisch, ??emait????ka, ????????????????????, ???????????????????? (??????????????????????), ??????????????????, Bamanankan, ???????????????, Brezhoneg, Bosanski, Catal??, M??ng-d????ng-ng?????, ??????????????, Cebuano, ?????????, ??esky, ?????????????????????? / ??????????????????????????????, ??????????????, Cymraeg, Dansk, Zazaki, ????????????????????, ????????????????, Emili??n e rumagn??l, English, Esperanto, Espa??ol, Eesti, Euskara, ??????????, Suomi, V??ro, F??royskt, Fran??ais, Arpetan, Furlan, Frysk, Gaeilge, ??????, G??idhlig, Galego, Ava??e'???, ?????????????????????, Gaelg, ??????????, ??????????????????, Fiji Hindi, Hrvatski, Krey??l ayisyen, Magyar, ??????????????, Interlingua, Bahasa Indonesia, Ilokano, Ido, ??slenska, Italiano, ?????????, Lojban, Basa Jawa, ?????????????????????, Kongo, Kalaallisut, ???????????????, ?????????, ????????????????-??????????????, Ripoarisch, Kurd??, ????????, Kernewek, ????????????????, Latina, Ladino, L??tzebuergesch, Limburgs, Ling??la, ?????????, Lietuvi??, Latvie??u, Basa Banyumasan, Malagasy, ????????????????????, ??????????????????, ???????????????, Bahasa Melayu, ????????????????, Nnapulitano, Nedersaksisch, ??????????????? ????????????, Nederlands, ???Norsk (nynorsk)???, ???Norsk (bokm??l)???, Nouormand, Din?? bizaad, Occitan, ????????????, Papiamentu, Deitsch, Norfuk / Pitkern, Polski, ????????????, ????????, Portugu??s, Runa Simi, Rumantsch, Romani, Rom??n??, ??????????????, ???????? ????????, Sardu, Sicilianu, Scots, S??megiella, Simple English, Sloven??ina, Sloven????ina, ???????????? / Srpski, Seeltersk, Svenska, Kiswahili, ???????????????, ??????????????????, ????????????, ?????????, T??rkmen??e, Tagalog, T??rk??e, ??????????????/Tatar??a, ????????????????????, ????????, Ti???ng Vi???t, Volap??k, Walon, Winaray, ??????, isiXhosa, ????????????, Yor??b??, Ze??uws, ??????, B??n-l??m-g??, ??????";
  
      function checkRecursively(map1, map2) {
        if (typeof map1 !== 'function' && typeof map2 !== 'function') {
          if (!map1 || typeof map1 !== 'object') {
            equal(map1, map2);
          } else {
            for (var key in map1) {
              checkRecursively(map1[key], map2[key]);
            }
          }
        }
      }
  
      test('Void', function() {
        equal(client.testVoid(), undefined);
      });
      test('Binary (String)', function() {
        var binary: string = '';
        for (var v = 255; v >= 0; --v) {
          binary += String.fromCharCode(v);
        }
        equal(client.testBinary(binary), binary);
      });
      test('Binary (Uint8Array)', function() {
        var binary: string = '';
        for (var v = 255; v >= 0; --v) {
          binary += String.fromCharCode(v);
        }
        var arr = new Uint8Array(binary.length);
        for (var i = 0; i < binary.length; ++i) {
          arr[i] = binary[i].charCodeAt(0);
        }
        const hexEncodedString = Array.from(arr, function(byte) {
            return String.fromCharCode(byte);
        }).join('')
        equal(client.testBinary(hexEncodedString), binary);
      });
      test('String', function() {
        equal(client.testString(''), '');
        equal(client.testString(stringTest), stringTest);
  
        var specialCharacters = 'quote: \" backslash:' +
          ' forwardslash-escaped: \/ ' +
            ' backspace: \b formfeed: \f newline: \n return: \r tab: ' +
              ' now-all-of-them-together: "\\\/\b\n\r\t' +
                ' now-a-bunch-of-junk: !@#$%&()(&%$#{}{}<><><';
                equal(client.testString(specialCharacters), specialCharacters);
      });
      test('Double', function() {
        equal(client.testDouble(0), 0);
        equal(client.testDouble(-1), -1);
        equal(client.testDouble(3.14), 3.14);
        equal(client.testDouble(Math.pow(2, 60)), Math.pow(2, 60));
      });
      test('Bool', function() {
        equal(client.testBool(true), true);
        equal(client.testBool(false), false);
      });
      test('I8', function() {
        equal(client.testByte(0), 0);
        equal(client.testByte(0x01), 0x01);
      });
      test('I32', function() {
        equal(client.testI32(0), 0);
        equal(client.testI32(Math.pow(2, 30)), Math.pow(2, 30));
        equal(client.testI32(-Math.pow(2, 30)), -Math.pow(2, 30));
      });
      test('I64', function() {
        equal(client.testI64(new Int64(0)), 0);
        equal(client.testI64(int64_2_pow_60), Math.pow(2, 52));
        equal(client.testI64(int64_minus_2_pow_60), -Math.pow(2, 52));
      });
  
      test('Struct', function() {
        var structTestInput: ThriftTest.Xtruct = new ThriftTest.Xtruct();
        structTestInput.string_thing = 'worked';
        structTestInput.byte_thing = 0x01;
        structTestInput.i32_thing = Math.pow(2, 30);
        structTestInput.i64_thing = int64_2_pow_60;
  
        var structTestOutput: ThriftTest.Xtruct = client.testStruct(structTestInput);
  
        equal(structTestOutput.string_thing, structTestInput.string_thing);
        equal(structTestOutput.byte_thing, structTestInput.byte_thing);
        equal(structTestOutput.i32_thing, structTestInput.i32_thing);
        equal(structTestOutput.i64_thing, structTestInput.i64_thing);
  
        equal(JSON.stringify(structTestOutput), JSON.stringify(structTestInput));
      });
  
      test('Nest', function() {
        var xtrTestInput: ThriftTest.Xtruct = new ThriftTest.Xtruct();
        xtrTestInput.string_thing = 'worked';
        xtrTestInput.byte_thing = 0x01;
        xtrTestInput.i32_thing = Math.pow(2, 30);
        xtrTestInput.i64_thing = int64_2_pow_60;
  
        var nestTestInput: ThriftTest.Xtruct2 = new ThriftTest.Xtruct2();
        nestTestInput.byte_thing = 0x02;
        nestTestInput.struct_thing = xtrTestInput;
        nestTestInput.i32_thing = Math.pow(2, 15);
  
        var nestTestOutput = client.testNest(nestTestInput);
  
        equal(nestTestOutput.byte_thing, nestTestInput.byte_thing);
        equal(nestTestOutput.struct_thing.string_thing, nestTestInput.struct_thing.string_thing);
        equal(nestTestOutput.struct_thing.byte_thing, nestTestInput.struct_thing.byte_thing);
        equal(nestTestOutput.struct_thing.i32_thing, nestTestInput.struct_thing.i32_thing);
        equal(nestTestOutput.struct_thing.i64_thing, nestTestInput.struct_thing.i64_thing);
        equal(nestTestOutput.i32_thing, nestTestInput.i32_thing);
  
        equal(JSON.stringify(nestTestOutput), JSON.stringify(nestTestInput));
      });
  
      test('Map', function() {
        var mapTestInput: {[k: number]: number;} = {7: 77, 8: 88, 9: 99};
  
        var mapTestOutput: {[k: number]: number;} = client.testMap(mapTestInput);
  
        for (var key in mapTestOutput) {
          equal(mapTestOutput[key], mapTestInput[key]);
        }
      });
  
      test('StringMap', function() {
        var mapTestInput: {[k: string]: string;} = {
          'a': '123', 'a b': 'with spaces ', 'same': 'same', '0': 'numeric key',
          'longValue': stringTest, stringTest: 'long key'
        };
  
        var mapTestOutput: {[k: string]: string;} = client.testStringMap(mapTestInput);
  
        for (var key in mapTestOutput) {
          equal(mapTestOutput[key], mapTestInput[key]);
        }
      });
  
      test('Set', function() {
        var setTestInput: number[] = [1, 2, 3];
        ok(client.testSet(setTestInput), setTestInput);
      });
  
      test('List', function() {
        var listTestInput: number[] = [1, 2, 3];
        ok(client.testList(listTestInput), listTestInput);
      });
  
      test('Enum', function() {
        equal(client.testEnum(ThriftTest.Numberz.ONE), ThriftTest.Numberz.ONE);
      });
  
      test('TypeDef', function() {
        equal(client.testTypedef(new Int64(69)), 69);
      });
  
      test('MapMap', function() {
        var mapMapTestExpectedResult: {[K: number]: {[k: number]: number}} = {
          '4': {'1': 1, '2': 2, '3': 3, '4': 4},
          '-4': {'-4': -4, '-3': -3, '-2': -2, '-1': -1}
        };
  
        var mapMapTestOutput = client.testMapMap(1);
  
  
        for (var key in mapMapTestOutput) {
          for (var key2 in mapMapTestOutput[key]) {
            equal(mapMapTestOutput[key][key2], mapMapTestExpectedResult[key][key2]);
          }
        }
  
        checkRecursively(mapMapTestOutput, mapMapTestExpectedResult);
      });
  
      test('Xception', function() {
        try {
          client.testException('Xception');
          ok(false, "expected an exception but there was no exception");
        } catch (e) {
          equal(e.errorCode, 1001);
          equal(e.message, 'Xception');
        }
      });
  
      test('no Exception', function() {
        try {
          client.testException('no Exception');
        } catch (e) {
          ok(false, "expected no exception but here was an exception");
        }
      });
  
      test('TException', function() {
        try {
          client.testException('TException');
          ok(false, "expected an exception but there was no exception");
        } catch (e) {
          ok(ok, "succesfully got exception");
        }
      });
  
      const crazy: ThriftTest.Insanity = {
        'userMap': { '5': new Int64(5), '8': new Int64(8) },
        'xtructs': [{
            'string_thing': 'Goodbye4',
            'byte_thing': 4,
            'i32_thing': 4,
            'i64_thing': new Int64(4)
        },
        {
            'string_thing': 'Hello2',
            'byte_thing': 2,
            'i32_thing': 2,
            'i64_thing': new Int64(2)
        }]
    };
      test('Insanity', function() {
        const insanity: {[k: number]: (ThriftTest.Insanity | {[k:number]: ThriftTest.Insanity})} = {
            '1': {
                '2': crazy,
                '3': crazy
            },
            '2': { '6': new ThriftTest.Insanity() }
        };
        var res = client.testInsanity(new ThriftTest.Insanity(crazy));
        ok(res, JSON.stringify(res));
        ok(insanity, JSON.stringify(insanity));
  
        checkRecursively(res, insanity);
      });
  
      console.log('------------------------------------------------------------');
      console.log('### All tests succeeded.');
      return 0;
    };
  
    try {
      var ret = execute();
      phantom.exit(ret);
    } catch (err) {
      // Catch all and exit to avoid hang.
      console.error(err);
      phantom.exit(1);
    }
  })();
  