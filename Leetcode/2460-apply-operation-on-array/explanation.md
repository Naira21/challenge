Розв'язок

1. Перебираємо дані nums. Рахуючи від першого, щоб не отримати ```panic: runtime error: index out of range```, порівнюємо, якщо поточний дорівнює попередньому, то попередній збільшии вдвічі, а поточний записати нулем. 

2. Додаємо змінну lastNonZeroFoundAt, щоб слідкувати за позицією, куди треба вставити наступне ненульове значення

3. Знову перебираємо nums
Ціль: ненульові значення винести вперед. 
Виконуємо: якщо значення за індексом не дорівнює нулю, то встановити першим у nums (nums[lastNonZeroFoundAt] = nums[0] ) значення nums[i]. 

У циклі перевіряти ще одну умову:
Якщо lastNonZeroFoundAt і i різні, це означає, що ми вже скопіювали значення nums[i] до позиції lastNonZeroFoundAt, і тепер можемо встановити значення в позиції i на 0.

Приклад для останього пункту
Припустимо, у нас є масив nums = [2, 0, 2, 4, 0, 8].

Після першого проходу (об'єднання однакових значень) масив може виглядати як [4, 0, 4, 0, 8, 0].
Другий прохід:
i = 0: nums[0] (4) переміщується на nums[0], lastNonZeroFoundAt збільшується на 1.
i = 1: nums[1] (0) не змінюється, lastNonZeroFoundAt не змінюється.
i = 2: nums[2] (4) переміщується на nums[1], і nums[2] стає 0.
i = 3: nums[3] (0) не змінюється, lastNonZeroFoundAt не змінюється.
i = 4: nums[4] (8) переміщується на nums[2], і nums[4] стає 0.
i = 5: nums[5] (0) не змінюється, lastNonZeroFoundAt не змінюється.
Результат: [4, 4, 8, 0, 0, 0].