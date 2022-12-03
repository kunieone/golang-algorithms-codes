#include <assert.h>
#include <stdio.h>
#include <stdlib.h>
#include <string.h>

void GetNext(char b[], int** next, int lenb) {
    // 首先我们吧next数组的第一个元素第二个元素初始化为-1和0；
    // 然后j直接从2下标开始遍历模式串
    (*next)[0] = -1;
    (*next)[1] = 0;
    int j = 2;
    int k = 0;

    while (j < lenb) {
        // 观察上图1，我们可以得到规律
        // 如果j-1处的字符和其对应的next数组处的字符相同，那么，j处的next数组元素就是k+1，也就是2
        // 在图2中，我们要求j位置的next值，而j-1处并不与K处的字符相等，那么我们接着遍历，
        // 让字符串继续回退到红k的next数组处代表的下标的字符处，绿色的K处字符为a，与j-1处字符相同，
        // 则next[j]=k+1；而此时的k为0，next[j]即为1

        // 从这两个图给的例子我们可以看到，如果j-1处的字符和next[j-1]处的字符相同，next[j]就等于k+1;
        // 如果不相同，则不断回退。直到找到相同的或者找不到时，K会被赋为-1，这时i++,j++，也就是主串和子串都往后移。而k会被接着赋为0；
        if (k == -1 || b[j - 1] == b[k]) {
            *((*next) + j) = k + 1;
            j++;
            k++;
        } else {
            k = *((*next) + k);
        }
    }
}

// 在这里实现KMP算法
int KMP(char a[], char b[]) {
    // 创建i变量控制主串a，j变量控制模式串b
    int i = 0;
    int j = 0;
    int lena = strlen(a);
    int lenb = strlen(b);

    // 动态创建next数组，用指针来操作
    int* next = (int*)malloc(lenb * sizeof(int));
    // int next[] = NewArray(lenb);
    assert(next);

    // 这里进入求next数组的步骤
    GetNext(b, &next, lenb);

    // 我们要确定这个KMP函数最后要返回的是-1还是另外的下标。
    // 在i小于主串长度和j小于模式串长度的时候进入这个循环，说明匹配过程在这个条件下要一直进行。
    // 如果i>lena了，就代表找不到匹配的串，如果j>lenb了，就代表找到了，就要返回对应下标。
    while (i < lena && j < lenb) {
        // 这里面的逻辑有些复杂，首先在a[i]==b[j]的情况下i++，j++大家肯定是理解的，这就代表匹配了，所以i++，j++。
        // 那么j==-1也进去是什么意思呢?
        // 首先我们要知道，j==-1其实是越界了的，然后在else语句中，j==next[j]，也就是说如果字符不匹配，j会调到模式串的next[j]处进行比对
        // 而如果知道j=0时字符仍然不匹配，j就跳到-1了（下图）。既然找不到匹配的字符，那么就将主串往后移，也即i++;模式串从0开始，即j++；
        if (j == -1 || a[i] == b[j]) {
            i++;
            j++;
        } else {
            j = next[j];
        }
    }
    // 如果模式串遍历完了，也就代表着找到了，返回主串下标
    if (j >= lenb)
        return i - j;

    free(next);
    next = NULL;
    // 没找到返回-1；
    return -1;
}

int main() {
    printf("%d", KMP("ababcabcdab", "abcd"));
    return 0;
}
