export default class AvatarColorCalculator {
    email_matrix: number[][] = []
    name_matrix: number[][] = []

    constructor(name : string | undefined, email : string | undefined) {
        const matrix_size = this._getMatrixSize(name!, email!)
        this.email_matrix = this._createMatrix(matrix_size, email!)
        this.name_matrix = this._createMatrix(matrix_size, name!)
    }

    public calculateColors(): {r: number, g: number, b: number} {
        let r = this._dot(this.email_matrix, this.name_matrix)
        let g = this._dot(this.email_matrix, this._transposition(this.name_matrix))
        return {r, g, b: (r + g) / 2}
    }

    private _transposition(matrix: Array<Array<number>>) : Array<Array<number>>  {
        let matrix_size = matrix.length
        let transposition_matrix : Array<Array<number>> = Array(matrix_size).fill(0).map(() => new Array(matrix_size).fill(0))
        for (let i = 0; i < matrix_size; i++) {
            for (let j = 0; j < matrix_size; j++) {
                transposition_matrix[i][j] = matrix[j][i]
            }
        }
        return transposition_matrix
    }  

    private _dot(matrix1: number[][], matrix2: number[][]) : number {
        let answer = 0
        for (let i = 0; i < matrix1.length; i++) {
            for (let j = 0; j < matrix1.length; j++) {
                answer = (answer + matrix1[i][j] * matrix2[i][j]) % 255
            }
        }

        return answer
    }

    private _createMatrix(n: number, word: string): number[][] {
        const matrix: number[][] = [];
        let index = 0;
    
        for (let i = 0; i < n; i++) {
          matrix[i] = [];
          for (let j = 0; j < n; j++) {
            if (index < word.length) {
              matrix[i][j] = word.charCodeAt(index);
              index++;
            } else {
              matrix[i][j] = 0;
            }
          }
        }
    
        return matrix;
    }

    private _getMatrixSize(word1: string, word2: string) : number {
        let word1_matrix_size = 0
        let word2_matrix_size = 0
        for (let i = 0; i < 100; i++) {
            if (i * i > word1.length) {
                word1_matrix_size = i
                break
            }
        }
        for (let i = 0; i < 100; i++) {
            if (i * i > word2.length) {
                word2_matrix_size = i
                break
            }
        }
        return Math.max(word1_matrix_size, word2_matrix_size)
    }
}