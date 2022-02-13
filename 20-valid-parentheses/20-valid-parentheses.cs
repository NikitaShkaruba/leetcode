public class Solution {
    private static Dictionary<char, char> matchingBracketsMap = new Dictionary<char, char>() {
      {'(', ')'},
      {'[', ']'},
      {'{', '}'},
    };

    private static bool IsBracket(char c) {
        return IsOpeningBracket(c) || IsClosingBracket(c);
    }

    private static bool IsOpeningBracket(char c) {
        return matchingBracketsMap.ContainsKey(c);	
    }

    private static bool IsClosingBracket(char c) {
        return matchingBracketsMap.ContainsValue(c);		
    }

    private static bool IsMatchingBracket(char openingBracket, char closingBracket) {
        return matchingBracketsMap[openingBracket] == closingBracket;
    }
  
    public bool IsValid(string str) {
        Stack<char> notYetClosedOpeningBrackets = new Stack<char>();

        foreach (char c in str) {
            if (!IsBracket(c)) {
                continue;
            }

            if (IsOpeningBracket(c)) {
              notYetClosedOpeningBrackets.Push(c);
                continue;
            }

            if (notYetClosedOpeningBrackets.Count == 0) {
                return false;
            }

            char lastPushedBracket = notYetClosedOpeningBrackets.Peek();
            if (!IsMatchingBracket(lastPushedBracket, c)) {
               return false;
            }

            notYetClosedOpeningBrackets.Pop();
        }

        return notYetClosedOpeningBrackets.Count == 0;
    }
}