// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

contract Voting {
    mapping(address => uint32) public votes;

    address[] candicate;

    mapping(address => bool) public isVote;

    //一个人只能给一个候选人投票
    function vote(address addr) public {
        if(!isVote[msg.sender]){
            votes[addr]+=1;
            candicate.push(addr);
        }
        isVote[msg.sender] = true;
      
    }

    function getVotesOf(address addr) view public returns (uint32){
        return votes[addr];
    }


    function resetVotes() public {
        //遍历数组清空票数
        for (uint i = 0; i < candicate.length; i++) {
            votes[candicate[i]] = 0;
        }
        isVote[msg.sender] = false;
    }
    
    function reverseString(string memory str) public pure returns(string memory){
        bytes memory bstr = bytes(str); // 将字符串转换成bytes数据
        uint256 len = bstr.length; //获取字节数
        for (uint i = 0; i < len / 2; i++) {
            bytes1 btempy = bstr[i];  
            bstr[i] = bstr[len - 1 - i];//交换两端字符
            bstr[len - 1 - i] = btempy; //赋值交换后的新字符
        }
        string memory str_new = string(bstr);
       return str_new;
    }


   // 罗马数字到值的映射
    mapping(bytes1 => uint256) private romanValues;
    
    constructor() {
        // 初始化罗马数字对应值
        romanValues['I'] = 1;
        romanValues['V'] = 5;
        romanValues['X'] = 10;
        romanValues['L'] = 50;
        romanValues['C'] = 100;
        romanValues['D'] = 500;
        romanValues['M'] = 1000;


        romanNumerals.push(RomanNumeral(1000, "M"));
        romanNumerals.push(RomanNumeral(900, "CM"));
        romanNumerals.push(RomanNumeral(500, "D"));
        romanNumerals.push(RomanNumeral(400, "CD"));
        romanNumerals.push(RomanNumeral(100, "C"));
        romanNumerals.push(RomanNumeral(90, "XC"));
        romanNumerals.push(RomanNumeral(50, "L"));
        romanNumerals.push(RomanNumeral(40, "XL"));
        romanNumerals.push(RomanNumeral(10, "X"));
        romanNumerals.push(RomanNumeral(9, "IX"));
        romanNumerals.push(RomanNumeral(5, "V"));
        romanNumerals.push(RomanNumeral(4, "IV"));
        romanNumerals.push(RomanNumeral(1, "I"));
    }
    
    // 主转换函数
    function romanToInt(string memory s) public view returns (uint256) {
        bytes memory roman = bytes(s);
        uint256 total = 0;
        uint256 prevValue = 0;
        
        // 从右向左处理罗马数字
        for (uint256 i = roman.length; i > 0; i--) {
            bytes1 currentChar = roman[i-1];
            require(isValidRoman(currentChar), "Invalid Roman numeral");
            
            uint256 currentValue = romanValues[currentChar];
            
            // 根据罗马数字规则处理
            if (currentValue < prevValue) {
                total -= currentValue;
            } else {
                total += currentValue;
            }
            
            prevValue = currentValue;
        }
        
        return total;
    }
    
    // 验证是否为有效罗马数字字符
    function isValidRoman(bytes1 c) private pure returns (bool) {
        return c == 'I' || c == 'V' || c == 'X' || c == 'L' || 
               c == 'C' || c == 'D' || c == 'M';
    }
    
    // 验证整个罗马数字字符串是否有效
    function isValidRomanNumeral(string memory s) public pure returns (bool) {
        bytes memory roman = bytes(s);
        if (roman.length == 0) return false;
        
        bytes1 lastChar;
        uint8 repeatCount = 1;
        
        for (uint256 i = 0; i < roman.length; i++) {
            bytes1 currentChar = roman[i];
            if (!isValidRoman(currentChar)) {
                return false;
            }
            
            // 检查重复规则
            if (currentChar == lastChar) {
                repeatCount++;
                // I, X, C, M最多重复3次；V, L, D不能重复
                if ((currentChar == 'I' || currentChar == 'X' || currentChar == 'C' || currentChar == 'M') && repeatCount > 3) {
                    return false;
                }
                if ((currentChar == 'V' || currentChar == 'L' || currentChar == 'D') && repeatCount > 1) {
                    return false;
                }
            } else {
                repeatCount = 1;
            }
            
            // 检查减法规则
            if (i > 0) {
                uint256 currentVal = getValue(currentChar);
                uint256 lastVal = getValue(lastChar);
                
                // 确保减法表示法有效
                if (currentVal > lastVal) {
                    bool validSubtraction = false;
                    if (lastChar == 'I' && (currentChar == 'V' || currentChar == 'X')) {
                        validSubtraction = true;
                    } else if (lastChar == 'X' && (currentChar == 'L' || currentChar == 'C')) {
                        validSubtraction = true;
                    } else if (lastChar == 'C' && (currentChar == 'D' || currentChar == 'M')) {
                        validSubtraction = true;
                    }
                    
                    if (!validSubtraction) {
                        return false;
                    }
                    
                    // 确保减法对不会重复出现（如IIV）
                    if (i > 1 && getValue(roman[i-2]) < currentVal) {
                        return false;
                    }
                }
            }
            
            lastChar = currentChar;
        }
        
        return true;
    }
    
    // 辅助函数：获取单个罗马数字的值
    function getValue(bytes1 c) private pure returns (uint256) {
        if (c == 'I') return 1;
        if (c == 'V') return 5;
        if (c == 'X') return 10;
        if (c == 'L') return 50;
        if (c == 'C') return 100;
        if (c == 'D') return 500;
        if (c == 'M') return 1000;
        revert("Invalid Roman numeral");
    }

    // 定义从整数到罗马数字的映射
    struct RomanNumeral {
        uint256 value;
        string symbol;
    }
    
    RomanNumeral[] private romanNumerals;
    

    
    // 主转换函数
    function intToRoman(uint256 num) public view returns (string memory) {
        require(num > 0 && num < 4000, "Number out of range (1-3999)");
        
        string memory roman = "";
        
        for (uint256 i = 0; i < romanNumerals.length; i++) {
            RomanNumeral memory current = romanNumerals[i];
            while (num >= current.value) {
                roman = string.concat(roman, current.symbol);
                num -= current.value;
            }
        }
        
        return roman;
    }
    
    
    //合并两个有序数组

    function mergeArrays(uint256[] memory arr1,uint256[] memory arr2) public pure returns (uint256[] memory){
        uint8 m = 0;
        uint256[] memory resultArray = new uint256[](arr1.length + arr2.length);
        for(uint i =0;i<arr1.length;i++){
            resultArray[m] = arr1[i];
            m++;
        }
         for(uint i =0;i<arr2.length;i++){
            resultArray[m] = arr2[i];
            m++;
        }
        //排序
        uint n = resultArray.length;
        for (uint i = 0; i < n - 1; i++) {
            for (uint j = 0; j < n - i - 1; j++) {
                if (resultArray[j] > resultArray[j + 1]) {
                    // 交换元素
                    (resultArray[j], resultArray[j + 1]) = (resultArray[j + 1], resultArray[j]);
                }
            }
        }
        return resultArray;
    }

    //在一个有序数组中查找目标值。
    function binarySearch(uint[] memory arr, uint target) public pure returns (bool) {
        uint left = 0;
        uint right = arr.length;
        
        while (left < right) {
            uint mid = left + (right - left) / 2;
            
            if (arr[mid] == target) {
                return true;
            } else if (arr[mid] < target) {
                left = mid + 1;
            } else {
                right = mid;
            }
        }
        
        return false;
    }

    
}