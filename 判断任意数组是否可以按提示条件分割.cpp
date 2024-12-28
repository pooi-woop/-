#include <iostream>
#include<vector>
using namespace std;
int main() {
    std::vector<int> vec;
    char num=0;
    std::cout << "Enter numbers (enter a e to stop): ";
    while (num != 'e') {
        std::cin >> num;
        vec.push_back((int)num);
    }
    /*
    for (unsigned int i = 0; i < vec.size(); ++i) {
        std::cout << vec[i] << " ";
    }
    std::cout << std::endl;
    cin.get();
    */
    bool canorcannot = 0;
        if (!(vec.size() % 2)) { canorcannot = 0; cout << canorcannot; return 0; }
        else{ canorcannot = 1;}
    int arrayN = vec.size();
    int* array = new int[vec.size()];
    for (arrayN = 0 ; arrayN <vec.size(); arrayN++) {
        array[arrayN] = vec[arrayN];
    }
    vector<int> count(100001, 0); 
    for (int num:vec) {
        count[num]++;
        if (count[num] > 2) {
            canorcannot = false;
            cout << canorcannot;
            return 0;
        }
    }
    cout << canorcannot << endl;
    return 0;
}